package main

import (
	"testing"
)

func Test_highlightFilename(t *testing.T) {
	s := "/home/vagrant/program_boot.pdf"
	actual := highlightFilename(s)
	expected := "<a target=\"_blank\" href=\"file:///home/vagrant/program_boot.pdf\"" + // Link
		">/home/vagrant/program_boot.pdf</a>" + // Text
		" <a href=\"file:///home/vagrant\" title=\"<< クリックでフォルダに移動\"><<</a>" //Directory
	if actual != expected {
		t.Fatalf("got: %v want: %v", actual, expected)
	}
}

func Test_highlightString(t *testing.T) {
	s := "/home/vagrant/Program/hoge3/program_boot.pdf"
	actual := highlightString(s, "program", "pdf")
	p := "<span style=\"background-color:#FFCC00;\">"
	q := "</span>"
	expected := "/home/vagrant/" +
		p + "Program" + q +
		"/hoge3/program_boot." +
		p + "pdf" + q
	if actual != expected {
		t.Fatalf("got: %v want: %v", actual, expected)
	}
}

func Test_andorPadding(t *testing.T) {
	for i, method := range []string{"and", "or"} {
		actual := andorPadding("this is test", method)
		expected := []string{
			"this.*is.*test", // AND Result
			"(this|is|test)", // OR Result
		}
		if actual != expected[i] {
			t.Fatalf("got: %v want: %v", actual, expected)
		}
	}
}

func Test_splitOutByte(t *testing.T) {
	b := []byte("hello\nmy\nname\n") // need last CRLF
	actual := splitOutByte(b)
	expected := []string{"hello", "my", "name"}
	for i, s := range actual {
		if s != expected[i] {
			t.Fatalf("got: %v want: %v", actual, expected)
		}
	}
}

func Test_htmlContents(t *testing.T) {
	test := []string{"/home/test/path", "this is a test word", "0", "1", "2", "3", "4", "5", "6", "7"}
	key := "is word"
	actual := htmlContents(test, key)
	expected := Result{
		Contents: []string{
			highlightFilename(test[0]),
			highlightString(test[1], "is", "word"),
		},
	}
	if actual.Contents[0] != expected.Contents[0] { // Filename test
		t.Fatalf("got: %v want: %v", actual.Contents[0], expected.Contents[0])
	}
	if actual.Contents[1] != expected.Contents[1] { // Contents test
		t.Fatalf("got: %v want: %v", actual.Contents[1], expected.Contents[1])
	}
	for i, e := range []string{"0", "1", "2", "3", "4", "5", "6", "7"} { // Stats test
		if actual.Stats[i] != e {
			t.Fatalf("got: %v want: %v", actual.Stats[i], e)
		}
	}
}

/*
func Test_htmlClause(t *testing.T) {
	// Case 1
	s := Search{Depth: "1", AndOr: "and"}
	actual := s.htmlClause()
	expected := `<!DOCTYPE html>
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charaset=utf-8">
			<title>Grep Server  </title>
			</head>
			  <body>
			    <form method="get" action="/search">
				  <!-- directory -->
				  <input type="text"
					  placeholder="検索対象フォルダのフルパスを入力してください(ex:/usr/bin ex:\\gr.net\ShareUsers\User\Personal)"
					  name="directory-path"
					  id="directory-path"
					  value=""
					  size="140"
					  title="検索対象フォルダのフルパスを入力してください(ex:/usr/bin ex:\\gr.net\ShareUsers\User\Personal)">
				  <a href=https://github.com/u1and0/grep-server/blob/master/README.md>Help</a>
				  <br>

				  <!-- file -->
				  <input type="text"
					  placeholder="検索キーワードをスペース区切りで入力してください"
					  name="query"
					  value=""
					  size="100"
					  title="検索キーワードをスペース区切りで入力してください">

				   <!-- depth -->
				   Lv
				   <select name="depth"
					  id="depth"
					  size="1"
					  title="Lv: 検索階層数を指定します。数字を増やすと検索速度は落ちますがマッチする可能性が上がります。">
					<option value="1" selected>1</option>
					<option value="2">2</option>
					<option value="3">3</option>
					<option value="4">4</option>
					<option value="5">5</option>
				  </select>
				 <!-- and/or -->
				 <input type="radio" value="and"
					title="スペース区切りをandとみなすかorとみなすか選択します"
					name="andor-search" checked="checked">and
					<input type="radio" value="or"
					title="スペース区切りをandとみなすかorとみなすか選択します"
					name="andor-search">or

				 <!-- encoding -->
				 <select name="encoding"
					id="encoding"
					size="1"
					title="文字エンコードを指定します。">
				  <option value="UTF-8" selected>utf-8</option>
					<option value="SHIFT-JIS">shift-jis</option>
					<option value="EUC-JP">euc-jp</option>
					<option value="ISO-2022-JP">iso-2022-jp</option>

				  </select>
				 <input type="submit" name="submit" value="Search">
			    </form>
				`
	if actual != expected {
		t.Fatalf("got: %v want: %v", actual, expected)
	}

	// Case 2
	s = Search{
		Keyword:  "test word",
		Path:     "/home/testuser",
		Depth:    "3",
		AndOr:    "and",
		Encoding: "shift-jis",
	}
	actual = s.htmlClause()
	expected = `<!DOCTYPE html>
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charaset=utf-8">
			<title>Grep Server test word /home/testuser</title>
			</head>
			  <body>
			    <form method="get" action="/search">
				  <!-- directory -->
				  <input type="text"
					  placeholder="検索対象フォルダのフルパスを入力してください(ex:/usr/bin ex:\\gr.net\ShareUsers\User\Personal)"
					  name="directory-path"
					  id="directory-path"
					  value="/home/testuser"
					  size="140"
					  title="検索対象フォルダのフルパスを入力してください(ex:/usr/bin ex:\\gr.net\ShareUsers\User\Personal)">
				  <a href=https://github.com/u1and0/grep-server/blob/master/README.md>Help</a>
				  <br>

				  <!-- file -->
				  <input type="text"
					  placeholder="検索キーワードをスペース区切りで入力してください"
					  name="query"
					  value="test word"
					  size="100"
					  title="検索キーワードをスペース区切りで入力してください">

				   <!-- depth -->
				   Lv
				   <select name="depth"
					  id="depth"
					  size="1"
					  title="Lv: 検索階層数を指定します。数字を増やすと検索速度は落ちますがマッチする可能性が上がります。">
					<option value="1">1</option>
					<option value="2">2</option>
					<option value="3" selected>3</option>
					<option value="4">4</option>
					<option value="5">5</option>
				  </select>

				 <!-- and/or -->
				 <input type="radio" value="and"
					title="スペース区切りをandとみなすかorとみなすか選択します"
					name="andor-search" checked="checked">and
					<input type="radio" value="or"
					title="スペース区切りをandとみなすかorとみなすか選択します"
					name="andor-search">or

				 <!-- encoding -->
				 <select name="encoding"
					id="encoding"
					size="1"
					title="文字エンコードを指定します。">
				  <option value="UTF-8">utf-8</option>
					<option value="SHIFT-JIS" selected>shift-jis</option>
					<option value="EUC-JP">euc-jp</option>
					<option value="ISO-2022-JP">iso-2022-jp</option>

				 <input type="submit" name="submit" value="Search">
			    </form>
				`
	if actual != expected {
		t.Fatalf("got: %v want: %v", actual, expected)
	}
}
*/
