#include <pistache/endpoint.h>

using namespace Pistache;

class HelloHandler : public Http::Handler
{
public:
    HTTP_PROTOTYPE(HelloHandler)

    void onRequest(const Http::Request& request, Http::ResponseWriter response)
    {
        response.send(Http::Code::Ok, "Hello, World\n");
    }
};

int main()
{
    Http::listenAndServe<HelloHandler>(Address(Ipv4::any(), Port(9000)));
}
