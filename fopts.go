
package fopts

type Option func(interface{})

type Builder struct {
    Opts []Option
}

func (b *Builder) With(opt Option) {
    b.Opts = append(b.Opts, opt)
}

func (b *Builder) Build(opts interface{}) {
    for _, opt := range b.Opts {
        opt(opts)
    }
}
