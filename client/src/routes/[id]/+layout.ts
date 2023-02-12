import { API } from "../../api";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ params }) => {
  let id = parseInt(params.id);
  let question = await API.getQuestion(id);

  return {
    question,
  };
};
