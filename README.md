Before we begin, let's just make it clear that you probably shouldn't build micro services. Micro services are a architectual style that might fit in some use cases, but before you're absolutely sure you know what you're building (think Netflix) it will be a lot easier to build a monolith and you will probably not run into scaling issues just yet. It is also a lot easier to split up a monolith that it is to clean up in a micro service mess


If you're looking for a client example: https://github.com/kimpettersen/svc-voicepayments

# svc-payments - an example server

Let's imagine you're building a product to make payments.
This is the micro service that deals with the moving money part of our product. It exposes four methods: https://github.com/kimpettersen/svc-payments/blob/master/proto/payments.proto#L11

The idea is that if you build a service like this, you can quickly build other service that won't have to solve the same hard problem again.

Next up we'll have a look at another micro service that are going to act as a client and call this service: https://github.com/kimpettersen/svc-voicepayments
