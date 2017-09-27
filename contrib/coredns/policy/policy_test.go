package policy

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	pdp "github.com/infobloxopen/themis/pdp-service"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

var (
	errFakePdp      = errors.New("fake PDP error")
	errFakeResolver = errors.New("fake Resolver error")
)

func TestPolicy(t *testing.T) {
	pm := PolicyPlugin{Next: handler(), options: make(map[uint16]edns0Map)}

	tests := []struct {
		query      string
		queryType  uint16
		response   *pdp.Response
		responseIP *pdp.Response
		errResp    error
		errRespIP  error
		status     int
		err        error
	}{
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			errResp:   errFakePdp,
			status:    dns.RcodeServerFailure,
			err:       errFakePdp,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_INDETERMINATE},
			status:    dns.RcodeRefused,
			err:       nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			errRespIP: errFakePdp,
			status:    dns.RcodeServerFailure,
			err:       errFakePdp,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_DENY},
			status:    dns.RcodeNameError,
			err:       nil,
		},
		{
			query:      "test.com.",
			queryType:  dns.TypeA,
			response:   &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_PERMIT},
			status:     dns.RcodeSuccess,
			err:        nil,
		},
		{
			query:      "test.com.",
			queryType:  dns.TypeA,
			response:   &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_DENY},
			status:     dns.RcodeNameError,
			err:        nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "221.228.88.194"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "redirect.biz"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "221.228.88.194"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "redirect.biz"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.org.",
			queryType: dns.TypeAAAA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "2001:db8:0:200:0:0:0:7"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "test.net"}}},
			status: dns.RcodeServerFailure,
			err:    errFakeResolver,
		},
		{
			query:     "test.net.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			status:    dns.RcodeServerFailure,
			err:       errFakeResolver,
		},
		{
			query:     "test.org.",
			queryType: dns.TypeAAAA,
			response:  &pdp.Response{Effect: pdp.Response_DENY},
			status:    dns.RcodeNameError,
			err:       nil,
		},
		{
			query:     "test.org.",
			queryType: dns.TypeAAAA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			status:    dns.RcodeSuccess,
			err:       nil,
		},
		{
			query:     "test.org.",
			queryType: dns.TypeAAAA,
			response: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"redirect_to", "string", "redirect.net"}}},
			status: dns.RcodeSuccess,
			err:    nil,
		},
		{
			query:     "test.org.",
			queryType: dns.TypeA,
			response: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"refuse", "string", "true"}}},
			status: dns.RcodeRefused,
			err:    nil,
		},
		{
			query:     "test.com.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			responseIP: &pdp.Response{Effect: pdp.Response_DENY,
				Obligation: []*pdp.Attribute{{"refuse", "string", "true"}}},
			status: dns.RcodeRefused,
			err:    nil,
		},
		{
			query:     "nxdomain.org.",
			queryType: dns.TypeA,
			response:  &pdp.Response{Effect: pdp.Response_PERMIT},
			status:    dns.RcodeNameError,
			err:       nil,
		},
	}

	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	for i, test := range tests {
		// Make request
		req := new(dns.Msg)
		req.SetQuestion(test.query, test.queryType)
		// Init test mock client
		pm.pdp = newTestClientInit(test.response, test.responseIP, test.errResp, test.errRespIP)
		// Handle request
		status, err := pm.ServeDNS(context.TODO(), rec, req)
		// Check status
		if test.status != status {
			t.Errorf("Case test[%d]: expected status %q but got %q\n", i, test.status, status)
		}
		// Check error
		if test.err != err {
			t.Errorf("Case test[%d]: expected error %v but got %v\n", i, test.err, err)
		}
	}

	pm.DebugSuffix = "debug."
	for i, test := range tests {
		// Make request
		req := new(dns.Msg)
		req.Question = make([]dns.Question, 1)
		req.Question[0] = dns.Question{test.query + pm.DebugSuffix, dns.TypeTXT, dns.ClassCHAOS}
		// Init test mock client
		pm.pdp = newTestClientInit(test.response, test.responseIP, test.errResp, test.errRespIP)
		// Handle request
		status, err := pm.ServeDNS(context.TODO(), rec, req)
		var test_status int
		var test_err error
		if test.err == errFakePdp {
			test_err = errFakePdp
			test_status = dns.RcodeServerFailure
		}
		if test.status == dns.RcodeRefused {
			test_status = dns.RcodeRefused
			test_err = nil
		}
		// Check status
		if test_status != status {
			t.Errorf("Case debug[%d]: expected status %q but got %q\n", i, test_status, status)
		}
		// Check error
		if test_err != err {
			t.Errorf("Case debug[%d]: expected error %v but got %v\n", i, test_err, err)
		}
	}
}

func handler() plugin.Handler {
	return plugin.HandlerFunc(func(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
		q := r.Question[0].Name
		switch q {
		case "test.com.":
			r.Answer = []dns.RR{
				test.A("test.com.		600	IN	A			10.240.0.1"),
			}
		case "test.org.":
			r.Answer = []dns.RR{
				test.AAAA("test.org.	600	IN	AAAA		21da:d3:0:2f3b:2aa:ff:fe28:9c5a"),
			}
		case "redirect.biz.":
			r.Answer = []dns.RR{
				test.A("redirect.biz.	600	IN	A			221.228.88.194"),
			}
		case "redirect.net.":
			r.Answer = []dns.RR{
				test.AAAA("redirect.net.	600	IN	AAAA		2001:db8:0:200:0:0:0:7"),
			}
		case "nxdomain.org.":
			w.WriteMsg(r)
			return dns.RcodeNameError, nil
		default:
			return dns.RcodeServerFailure, errFakeResolver
		}
		w.WriteMsg(r)
		return dns.RcodeSuccess, nil
	})
}

type testClient struct {
	nextResponse   *pdp.Response
	nextResponseIP *pdp.Response
	errResponse    error
	errResponseIP  error
}

func newTestClientInit(nextResponse *pdp.Response, nextResponseIP *pdp.Response,
	errResponse error, errResponseIP error) *testClient {
	return &testClient{
		nextResponse:   nextResponse,
		nextResponseIP: nextResponseIP,
		errResponse:    errResponse,
		errResponseIP:  errResponseIP,
	}
}

func (c *testClient) Connect() error { return nil }
func (c *testClient) Close()         {}
func (c *testClient) Validate(ctx context.Context, in, out interface{}) error {
	if in != nil {
		p := in.(pdp.Request)
		for _, a := range p.Attributes {
			if a.Id == "address" {
				if c.errResponseIP != nil {
					return c.errResponseIP
				}
				if c.nextResponseIP != nil {
					return fillResponse(c.nextResponseIP, out)
				}
				continue
			}
		}
	}
	if c.errResponse != nil {
		return c.errResponse
	}
	return fillResponse(c.nextResponse, out)
}

func (c *testClient) ModalValidate(in, out interface{}) error {
	return c.Validate(context.Background(), in, out)
}

func fillResponse(in *pdp.Response, out interface{}) error {
	r, ok := out.(*pdp.Response)
	if !ok {
		return fmt.Errorf("testClient can only translate response to *pb.Response type but got %T", out)
	}
	r.Effect = in.Effect
	r.Obligation = in.Obligation
	return nil
}

func makeRequestWithEDNS0(code uint16, hexstring string, nonlocal bool) *dns.Msg {
	req := new(dns.Msg)
	o := new(dns.OPT)
	o.Hdr.Name = "."
	o.Hdr.Rrtype = dns.TypeOPT
	if nonlocal {
		e := new(dns.EDNS0_SUBNET)
		o.Option = append(o.Option, e)
	} else {
		e := new(dns.EDNS0_LOCAL)
		e.Code = code
		src := []byte(hexstring)
		dst := make([]byte, hex.DecodedLen(len(src)))
		hex.Decode(dst, src)
		e.Data = dst
		o.Option = append(o.Option, e)
	}
	req.Extra = append(req.Extra, o)
	return req
}

func TestEdns(t *testing.T) {
	pm := PolicyPlugin{options: make(map[uint16]edns0Map)}

	// Add EDNS mapping
	if err := pm.AddEDNS0Map("0xfffa", "client_id", "hex", "string", "0", "32"); err != nil {
		t.Errorf("Expected error 'nil' but got %v\n", err)
	}
	if err := pm.AddEDNS0Map("0xfffb", "source_ip", "address", "address", "0", "0"); err != nil {
		t.Errorf("Expected error 'nil' but got %v\n", err)
	}
	if err := pm.AddEDNS0Map("0xfffc", "client_name", "bytes", "string", "0", "0"); err != nil {
		t.Errorf("Expected error 'nil' but got %v\n", err)
	}

	tests := []struct {
		name     string
		code     uint16
		data     string
		ip       string
		attr     map[string]*pdp.Attribute
		nonlocal bool
	}{
		{
			name: "Test different than EDNS0_LOCAL option",
			ip:   "192.168.0.1",
			attr: map[string]*pdp.Attribute{
				"source_ip": {Id: "source_ip", Type: "address", Value: "192.168.0.1"},
			},
			nonlocal: true,
		},
		{
			name: "Test option that not in config mapping",
			code: 0xfffd,
			data: "cafecafe",
			ip:   "192.168.0.2",
			attr: map[string]*pdp.Attribute{
				"source_ip": {Id: "source_ip", Type: "address", Value: "192.168.0.2"},
			},
		},
		{
			name: "Test option handled as hex",
			code: 0xfffa,
			data: "4e7e318384088e7d4f3dbc96219ee5d4",
			ip:   "192.168.0.3",
			attr: map[string]*pdp.Attribute{
				"source_ip": {Id: "source_ip", Type: "address", Value: "192.168.0.3"},
				"client_id": {Id: "client_id", Type: "string", Value: "4e7e318384088e7d4f3dbc96219ee5d4"},
			},
		},
		{
			name: "Test option 'source_ip' handled as address",
			code: 0xfffb,
			data: "aca80001", // 172.168.0.1 in hex
			ip:   "192.168.0.4",
			attr: map[string]*pdp.Attribute{
				"source_ip":       {Id: "source_ip", Type: "address", Value: "172.168.0.1"},
				"proxy_source_ip": {Id: "proxy_source_ip", Type: "address", Value: "192.168.0.4"},
			},
		},
		{
			name: "Test option handled as bytes",
			code: 0xfffc,
			data: "637573746f6d6572", // "customer" in hex
			ip:   "192.168.0.5",
			attr: map[string]*pdp.Attribute{
				"source_ip":   {Id: "source_ip", Type: "address", Value: "192.168.0.5"},
				"client_name": {Id: "client_name", Type: "string", Value: "customer"},
			},
		},
	}

	for _, test := range tests {
		req := makeRequestWithEDNS0(test.code, test.data, test.nonlocal)
		attr := pm.getAttrsFromEDNS0(req, test.ip)
		mapAttr := make(map[string]*pdp.Attribute)
		for _, a := range attr {
			mapAttr[a.Id] = a
		}
		if len(attr) != len(mapAttr) {
			t.Errorf("%s: array %q transformed to map %q\n", test.name, attr, mapAttr)
		}
		if !reflect.DeepEqual(test.attr, mapAttr) {
			t.Errorf("%s: expected attributes %q but got %q\n", test.name, test.attr, mapAttr)
		}
	}
}
