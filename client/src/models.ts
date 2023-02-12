export interface IAnswer {
  id: number;
  answer: string;
  votes: number;
  questionId: number;
}

export interface IQuestion {
  id: number;
  question: string;
  answers: IAnswer[];
}
