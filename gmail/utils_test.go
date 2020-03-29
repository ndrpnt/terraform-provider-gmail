package gmail

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Test_ignoreLeadingTrailingRepeatedSpaces(t *testing.T) {
	t.Parallel()
	type args struct {
		k   string
		old string
		new string
		d   *schema.ResourceData
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				old: "foobar",
				new: "foobar",
			},
			want: true,
		},
		{
			args: args{
				old: "foo bar",
				new: "foo bar",
			},
			want: true,
		},
		{
			args: args{
				old: "foo  bar",
				new: "foo  bar",
			},
			want: true,
		},
		{
			args: args{
				old: "  foobar",
				new: "  foobar",
			},
			want: true,
		},
		{
			args: args{
				old: "foobar  ",
				new: "foobar  ",
			},
			want: true,
		},
		{
			args: args{
				old: "   foo   bar   ",
				new: "   foo   bar   ",
			},
			want: true,
		},
		{
			args: args{
				old: "foo  bar",
				new: "foo bar",
			},
			want: true,
		},
		{
			args: args{
				old: " foobar",
				new: "foobar",
			},
			want: true,
		},
		{
			args: args{
				old: "foobar ",
				new: "foobar",
			},
			want: true,
		},
		{
			args: args{
				old: "   foo   bar   ",
				new: "foo bar",
			},
			want: true,
		},
		{
			args: args{
				old: "   foo bar   ",
				new: "foo   bar",
			},
			want: true,
		},
		{
			args: args{
				old: "   foo   bar   baz   ",
				new: "foo bar baz",
			},
			want: true,
		},
		{
			args: args{
				old: "\t\n\v\f\r \u0085\u00A0foo\t\n\v\f\r \u0085\u00A0bar\t\n\v\f\r \u0085\u00A0baz\t\n\v\f\r \u0085\u00A0",
				new: "foo bar baz",
			},
			want: true,
		},
		{
			args: args{
				old: "foobar",
				new: "foobarbaz",
			},
			want: false,
		},
		{
			args: args{
				old: "foo bar",
				new: "foobar",
			},
			want: false,
		},
		{
			args: args{
				old: " foo bar ",
				new: "foobar",
			},
			want: false,
		},
		{
			args: args{
				old: "foo\u0085bar",
				new: "foo bar",
			},
			want: true,
		},
		{
			args: args{
				old: "foo_bar",
				new: "foo bar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		var name string
		if tt.want {
			name = fmt.Sprintf("should consider '%s' and '%s' equivalent", tt.args.old, tt.args.new)
		} else {
			name = fmt.Sprintf("should consider '%s' and '%s' different", tt.args.old, tt.args.new)
		}
		t.Run(name, func(t *testing.T) {
			if got := ignoreLeadingTrailingRepeatedSpaces(tt.args.k, tt.args.old, tt.args.new, tt.args.d); got != tt.want {
				t.Errorf("ignoreLeadingTrailingRepeatedSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noLeadingTrailingRepeatedSpaces(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input      string
		wantErrors bool
	}{
		{
			input:      "foo",
			wantErrors: false,
		},
		{
			input:      " foo",
			wantErrors: true,
		},
		{
			input:      "foo ",
			wantErrors: true,
		},
		{
			input:      "foo bar",
			wantErrors: false,
		},
		{
			input:      "foo  bar",
			wantErrors: true,
		},
		{
			input:      "foo bar baz",
			wantErrors: false,
		},
		{
			input:      "foo bar  baz",
			wantErrors: true,
		},
		{
			input:      "foo bar baz  ",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\t",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\n",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\v",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\f",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\r",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\u00A0",
			wantErrors: true,
		},
		{
			input:      "foo bar baz\u0085",
			wantErrors: true,
		},
	}
	for _, tt := range tests {
		var name string
		if tt.wantErrors {
			name = fmt.Sprintf("should yield errors with input '%s'", tt.input)
		} else {
			name = fmt.Sprintf("should not yield errors with input '%s'", tt.input)
		}
		t.Run(name, func(t *testing.T) {
			_, errs := noLeadingTrailingRepeatedSpaces()(tt.input, "dummy_key")
			if tt.wantErrors && len(errs) == 0 {
				t.Errorf("noLeadingTrailingRepeatedSpaces() errors = none, want some")
			} else if !tt.wantErrors && len(errs) > 0 {
				t.Errorf("noLeadingTrailingRepeatedSpaces() errors = %v, want none", errs)
			}
		})
	}
}
