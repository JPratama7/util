package token

import "time"

type Generator[ID ~string, T IdGetter[ID]] struct {
	public, private string
	duration        time.Duration
}

func NewGenerator[ID ~string, T IdGetter[ID]](public, private string, duration time.Duration) *Generator[ID, T] {
	return &Generator[ID, T]{
		public:   public,
		private:  private,
		duration: duration,
	}
}
