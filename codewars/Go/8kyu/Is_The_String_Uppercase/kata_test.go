// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "../Is_The_String_Uppercase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values", func() {
		Expect(MyString("a").IsUpperCase()).To(Equal(false))
		Expect(MyString("b").IsUpperCase()).To(Equal(false))
		Expect(MyString("c").IsUpperCase()).To(Equal(false))
		Expect(MyString("d").IsUpperCase()).To(Equal(false))
		Expect(MyString("e").IsUpperCase()).To(Equal(false))
		Expect(MyString("f").IsUpperCase()).To(Equal(false))
		Expect(MyString("g").IsUpperCase()).To(Equal(false))
		Expect(MyString("h").IsUpperCase()).To(Equal(false))
		Expect(MyString("i").IsUpperCase()).To(Equal(false))
		Expect(MyString("j").IsUpperCase()).To(Equal(false))
		Expect(MyString("k").IsUpperCase()).To(Equal(false))
		Expect(MyString("l").IsUpperCase()).To(Equal(false))
		Expect(MyString("m").IsUpperCase()).To(Equal(false))
		Expect(MyString("n").IsUpperCase()).To(Equal(false))
		Expect(MyString("o").IsUpperCase()).To(Equal(false))
		Expect(MyString("p").IsUpperCase()).To(Equal(false))
		Expect(MyString("q").IsUpperCase()).To(Equal(false))
		Expect(MyString("r").IsUpperCase()).To(Equal(false))
		Expect(MyString("s").IsUpperCase()).To(Equal(false))
		Expect(MyString("t").IsUpperCase()).To(Equal(false))
		Expect(MyString("u").IsUpperCase()).To(Equal(false))
		Expect(MyString("v").IsUpperCase()).To(Equal(false))
		Expect(MyString("w").IsUpperCase()).To(Equal(false))
		Expect(MyString("x").IsUpperCase()).To(Equal(false))
		Expect(MyString("y").IsUpperCase()).To(Equal(false))
		Expect(MyString("z").IsUpperCase()).To(Equal(false))
		Expect(MyString("abcdefghijklmnopqrstuvwxyz").IsUpperCase()).To(Equal(false))
		Expect(MyString("ABCDEFGHIJKLMNOPQRSTUVWXYz").IsUpperCase()).To(Equal(false))
		Expect(MyString("false").IsUpperCase()).To(Equal(false))
		Expect(MyString("true").IsUpperCase()).To(Equal(false))
		Expect(MyString("False").IsUpperCase()).To(Equal(false))
		Expect(MyString("True").IsUpperCase()).To(Equal(false))

		Expect(MyString("WHAT DOES THE FOX SAY").IsUpperCase()).To(Equal(true))
		Expect(MyString("HTML CSS JAVASCRIPT PYTHON C PERL LISP JAVA GO RUBY NODEJS RUST SCALA").IsUpperCase()).To(Equal(true))
	})
})
