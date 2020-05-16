package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/beacon/deployer/pkg/config"
	"google.golang.org/grpc"

	pb "github.com/beacon/deployer/pkg/proto"
)

var cfg = &config.Config{
	Addr: ":9000",
	TLS:  &config.TLSConfig{},
}

func init() {
	var cert = `-----BEGIN CERTIFICATE-----
MIICjDCCAfWgAwIBAgIBATANBgkqhkiG9w0BAQsFADBXMQswCQYDVQQGEwJDTjER
MA8GA1UECAwIU2hhbmdoYWkxETAPBgNVBAcMCFNoYW5naGFpMQ4wDAYDVQQKDAVF
dGhhbjESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIwMDUxNTA3NDIwOVoXDTIxMDUx
NTA3NDIwOVowRDELMAkGA1UEBhMCQ04xETAPBgNVBAgMCFNoYW5naGFpMQ4wDAYD
VQQKDAVFdGhhbjESMBAGA1UEAwwJbG9jYWxob3N0MIGfMA0GCSqGSIb3DQEBAQUA
A4GNADCBiQKBgQDa7CkVjL4jKJiKEDa5LQKyjrSXFrpKgRdJXcu4grUjxIxUt5LJ
oycMn+5d9MFDgUJ8OoZm+TP34OCSANCXxhjeA+Y25rxSOQJYEdFsi2jiLw2lJBx0
fgFXTnWVOvZPQnihy0KKIqxOFxQqC2aRuIMj+e2IWFYVSa+5M+YnMDl9fwIDAQAB
o3sweTAJBgNVHRMEAjAAMCwGCWCGSAGG+EIBDQQfFh1PcGVuU1NMIEdlbmVyYXRl
ZCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQUJcXCtS6K2fwQc/0IOfo1llRzQ2QwHwYD
VR0jBBgwFoAUHsnvuVp9j2bKOYqkLJ4Yx8qkxIIwDQYJKoZIhvcNAQELBQADgYEA
pPd6qe6bImsk4lB9JCDtjwr8KydbNCTtjAn4bqTHZaOllgvsMX+Wa5xSWkMDGoBh
n742uXrU8os0K/6YF/HNkn1lTWcnBj9udkuAPu3jOxC3PuhB29ki5cROjPfeJOIr
zHAmofkZA/MFhLPaYEp7MTsufSWXJluF6DW2IsP/pRA=
-----END CERTIFICATE-----
`
	var key = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDa7CkVjL4jKJiKEDa5LQKyjrSXFrpKgRdJXcu4grUjxIxUt5LJ
oycMn+5d9MFDgUJ8OoZm+TP34OCSANCXxhjeA+Y25rxSOQJYEdFsi2jiLw2lJBx0
fgFXTnWVOvZPQnihy0KKIqxOFxQqC2aRuIMj+e2IWFYVSa+5M+YnMDl9fwIDAQAB
AoGBAKwKFqc2uV0L1AP7heWzt4D2oHhzheZy2Afxy9yt6we2t8kDkjkfG9rR/BKd
+xq/g634mBJoGCzd7d0PAt2i8XpjAd7K96qIwvtYE8K64LPfniET9vrGUIFyRgDO
zdvjTao/Q3XqInBG4nR2CcXFw0YLXkWxpTjOoas0iK+FPolBAkEA98gs1ayFq8im
Q0silbnMftJHj1rQE40iVLl+RM+IDg5X+8a2LipTVz6YPIqF3L/hpzEXgEB93M+F
GDC6w4f2UQJBAOIu81iyZwS2hxfn/RR5ot/AaEcLZL6zWwwMqLe/CV5hQPzltSED
cZyUG+QfesGzg06CavB3SHOsXr7oGCl1ss8CQQCVcI9J16FO99b+4wPa1ZI2MtCh
7x1rjUVVYAJ9scTW5WO/IBukQDa/easLaGhPuRJ5aaxI15yRXj9hVZJud0PhAkB/
3j8gc6sd1PrGnxZKTwGvMR1CnMRVsxvT0gxH5K4tNxoAXvRpN4MxG+Iws0M44n1n
Ev/V9fl1u4rMrnWKasmtAkEAurWnsfEIbisinIoiRJpjypp9BP98rF/dc/t8tuTu
KxVEiITF3MPcsthztvY5RGvmCo9/T9vDVlJEGSF04F1B3A==
-----END RSA PRIVATE KEY-----
`
	cfg.TLS.CertFile = path.Join(os.TempDir(), "cert")
	cfg.TLS.KeyFile = path.Join(os.TempDir(), "key")
	ioutil.WriteFile(cfg.TLS.CertFile, []byte(cert), 0644)
	ioutil.WriteFile(cfg.TLS.KeyFile, []byte(key), 0644)
}

func TestPlainServer(t *testing.T) {
	cfg2 := &config.Config{
		Addr: cfg.Addr,
	}
	srv := New(cfg2)
	go func() {
		if err := srv.ListenAndServe(cfg2); err != nil && err != http.ErrServerClosed {
			t.Log("Server stopped with error:", err)
		}
	}()
	time.Sleep(time.Second)

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost"+cfg2.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateDeployStatus(ctx, &pb.DeployStatus{
		Id: "unittest",
	})
	if err != nil {
		log.Fatalf("could not connect grpc: %v", err)
	}
	log.Printf("Reply: %s", r.Message)

	resp, err := http.Post("http://localhost"+cfg2.Addr+"/actions", "applications/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
	srv.Shutdown()
}

func TestTLSServer(t *testing.T) {
	srv := New(cfg)
	go func() {
		if err := srv.ListenAndServe(cfg); err != nil && err != http.ErrServerClosed {
			t.Log("Server stopped with error:", err)
		}
	}()
	time.Sleep(time.Second)
	sysPool, err := x509.SystemCertPool()
	if err != nil {
		t.Fatalf("failed to load certificate pool from system: %v", err)
	}
	var ca = `-----BEGIN CERTIFICATE-----
MIICjDCCAfWgAwIBAgIUZ45uACTc8223LmVXyTqLAc0FLOkwDQYJKoZIhvcNAQEL
BQAwVzELMAkGA1UEBhMCQ04xETAPBgNVBAgMCFNoYW5naGFpMREwDwYDVQQHDAhT
aGFuZ2hhaTEOMAwGA1UECgwFRXRoYW4xEjAQBgNVBAMMCWxvY2FsaG9zdDAgFw0y
MDA1MTUwNzM3MTZaGA8yMjk0MDIyODA3MzcxNlowVzELMAkGA1UEBhMCQ04xETAP
BgNVBAgMCFNoYW5naGFpMREwDwYDVQQHDAhTaGFuZ2hhaTEOMAwGA1UECgwFRXRo
YW4xEjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkC
gYEA37g2sBFGyoKWA+LmFGJGOe15FSzfxtVejcsR/CpSiXRvizNX08Yz5TMAC3bW
Yw6ZPxLMbEl4/cL/woFKY9Vuinvj04zq77yszjc05kPwb4tFz7P4IxG48serNdV1
0n+9HiCKbasXa8Y4MFGjOohFik3UtQAFi/oPnPYuQk+EDvcCAwEAAaNTMFEwHQYD
VR0OBBYEFB7J77lafY9myjmKpCyeGMfKpMSCMB8GA1UdIwQYMBaAFB7J77lafY9m
yjmKpCyeGMfKpMSCMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEA
WmXe2W6GdOEwyM1kwX90+lwiB2W725GrfRrPUHkLJnrlvGpAnz0p4W7Jeoipm+cp
CB6e1BOf3UQL9JCnOtcXYWasw5uN2WWpa1BhbZKw+lD2luLImYcVxp9oPgObTpYB
Iz4xPm1u+03+NA805NyA+Q9KAFv/62Gz/qQAFSmTkfs=
-----END CERTIFICATE-----
`
	if ok := sysPool.AppendCertsFromPEM([]byte(ca)); !ok {
		t.Fatal("failed to add ca")
	}

	tlsConfig := &tls.Config{
		RootCAs: sysPool,
	}
	cert, err := tls.LoadX509KeyPair(cfg.TLS.CertFile, cfg.TLS.KeyFile)
	if err != nil {
		t.Fatal(err)
	}
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

	creds := credentials.NewTLS(tlsConfig)
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost"+cfg.Addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateDeployStatus(ctx, &pb.DeployStatus{
		Id: "unittest",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.Message)

	addr := "https://localhost" + cfg.Addr + "/actions"
	t.Log(addr)
	req, _ := http.NewRequest("POST", addr, nil)
	req.Header.Set("Content-Type", "applications/json")
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := httpClient.Post(addr, "applications/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
	srv.Shutdown()
}
