package helpers

func IsConnectionError(err error) bool {
	return err.Error() == "dial tcp: lookup" || err.Error() == "getaddrinfow: The requested name is valid, but no data of the requested type was found"
}

func IsUnexpectedEOF(err error) bool {
	return err.Error() == "EOF" || err.Error() == "unexpected EOF"
}
