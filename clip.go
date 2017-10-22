// +build !windows
// +build !darwin

package clip

type Clip struct {
	n   int
	err error
}

func New() (*Clip, error) {
	return &Clip{}, nil
}
func (c *Clip) Read(p []byte) (n int, err error) {
	return c.n, c.err
}

func (c *Clip) Write(p []byte) (n int, err error) {
	return c.n, c.err
}

func (c *Clip) Close() (err error) {
	return c.err
}
