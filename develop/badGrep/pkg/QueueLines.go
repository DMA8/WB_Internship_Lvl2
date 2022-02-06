package pkg

type QueueLines struct {
	CurrentString	string
	BeforeStrings	[]string
	AfterStrings	[]string
	NLinesBefore	int
	NLinesAfter		int
}

func NewQueueLines(nLineAfter, nLineBefore int) *QueueLines {
	ql := QueueLines{}
	ql.NLinesAfter = nLineAfter
	ql.NLinesBefore = nLineBefore
	ql.AfterStrings = make([]string, 0, ql.NLinesAfter)
	ql.BeforeStrings = make([]string, 0, ql.NLinesBefore)
	return &ql
}

type ManyStringsReader interface {
	ReadNString(n int) []string
}

func (q *QueueLines)AddStringAfter(inpStr string) {
	if len(q.AfterStrings) < q.NLinesAfter {
		q.AfterStrings = append(q.AfterStrings, inpStr)
	} else {
		q.AfterStrings = q.AfterStrings[1:]
		q.AfterStrings = append(q.AfterStrings, inpStr)
	}
}

func (q *QueueLines)AddStringBefore(inpStr string) {
	if len(q.BeforeStrings) < q.NLinesBefore {
		q.BeforeStrings = append(q.BeforeStrings, inpStr)
	} else {
		q.BeforeStrings = q.BeforeStrings[1:]
		q.BeforeStrings = append(q.BeforeStrings, inpStr)
	}
}

func (q *QueueLines)FirstFillAfter(reader ManyStringsReader) {
	q.AfterStrings = reader.ReadNString(q.NLinesAfter)
}
