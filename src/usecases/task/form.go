package task

type RequestForm struct {
	UserId      int    `schema:"userId,required"`
	Description string `schema:"description,required"`
	LawyerId    int    `schema:"lawyerId,required"`
}

type AnnounceForm struct {
	UserId      int    `schema:"userId,required"`
	Description string `schema:"description,required"`
}
