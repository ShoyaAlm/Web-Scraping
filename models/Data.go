package models

type Data struct {
	lines      []interface{} // could be paragraphs, headings or links
	photos     []interface{} // pics & stuff
	webFormat  []interface{} // JSON or XML format
	webDocs    []interface{} // PDFs, Docs, Word files
	multiMedia []interface{} // videos, audio

}
