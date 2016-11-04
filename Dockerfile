FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-auth /kyp-auth
ENV KYP_AUTH_DB="postgres://postgres:@docker/kyp_auth_dev?sslmode=disable"
ENV KYP_AUTH_PORT="3000"
ENV KYP_SECRET_KEY="c91267c27a8599ca0480ea505487d052e3b63a1dd39819db853225a518200399"
ENTRYPOINT ["/kyp-auth"]
