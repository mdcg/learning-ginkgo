package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mdcg/learning-ginkgo/books"
)

var _ = Describe("Book", func() {
	var (
		longBook  books.Book
		shortBook books.Book
		book      books.Book
		err       error
		json      string
	)

	BeforeEach(func() {
		longBook = books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = books.Book{
			Title:  "Fox in Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}

		json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
	})

	JustBeforeEach(func() {
		book, err = books.NewBookFromJSON(json)
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				json = `{
					"title":"Les Miserables",
					"author":"Victor Hugo",
					"pages":1488ops
				}`
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})
