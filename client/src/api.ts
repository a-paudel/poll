import { browser, dev } from "$app/environment";
import type { IQuestion } from "./models";

let BASE_URL = "http://localhost:8000";
let WS_URL = "ws://localhost:8000";

if (!dev && browser) {
  BASE_URL = window.location.origin;
  WS_URL = window.location.origin.replace("https://", "wss://");
  WS_URL = window.location.origin.replace("http://", "ws://");
}

async function getQuestion(id: number): Promise<IQuestion> {
  const resp = await fetch(`${BASE_URL}/api/questions/${id}`);
  let data = await resp.json();
  return data;
}

function listenQuestion(id: number): WebSocket {
  let ws = new WebSocket(`${WS_URL}/api/questions/${id}/listen`);
  return ws;
}

async function createQuestion(input: {
  question: string;
  answers: string[];
}): Promise<IQuestion> {
  const resp = await fetch(`${BASE_URL}/api/questions`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(input),
  });
  let data = await resp.json();
  return data;
}

async function voteQuestion(id: number, answerId: number): Promise<IQuestion> {
  const resp = await fetch(`${BASE_URL}/api/questions/${id}/vote`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ answerId }),
  });
  return await resp.json();
}

export const API = {
  createQuestion,
  getQuestion,
  listenQuestion,
  voteQuestion,
};
