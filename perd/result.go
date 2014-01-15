package perd

type Result interface {
  Bytes() []byte
}

func NewResult(out, err []byte, code int) Result {
  return &result{out, err, code}
}

type result struct {
  stdOut      []byte
  stdErr      []byte
  statusCode  int
}

func (r *result) Bytes () []byte {
  return []byte(r.stdOut)
}
