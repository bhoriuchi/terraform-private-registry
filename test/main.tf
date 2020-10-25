terraform {
    required_providers {
        foo = {
            version = ">=2.1.0"
            source = "localhost:8443/acme/foo"
        }
    }
}

resource "foo" "test" {
    content     = "foo!"
    filename = "${path.module}/foo.bar"
}
