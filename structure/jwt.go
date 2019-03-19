package structure

/*
	HEADER : ALGORITHM & TOKEN TYPE
	{
		"alg": "HS256",
		"typ": "JWT"
	}
	PAYLOAD : DATA
	{
		"firstname" : "apisit"
	}
	VERIFY SIGNATURE
	{
		HMACSHA256(
			base64UrlEncode(HEADER) + "."  +
			base64UrlEncode(PAYLOAD) + "." +
			base64UrlEncode("SIGNATURE")
		)
	}

*/
type JWT struct {
	HEADER    string
	PAYLOAD   string
	SIGNATURE string
}
