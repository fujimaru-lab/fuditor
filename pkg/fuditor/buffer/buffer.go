package buffer

// Tab タブ
type Tab struct {
	Tabs  []Text
	TabNo int
}

// Text タブ内のテキスト
type Text struct {
	Lines map[int]Line
}

// Line テキスト内の行ごとのテキストデータ
type Line struct {
	Line map[int][]rune
}
