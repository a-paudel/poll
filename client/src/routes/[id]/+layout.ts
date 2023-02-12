import { goto } from "$app/navigation";
import Swal from "sweetalert2";
import { API } from "../../api";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ params }) => {
  let id = parseInt(params.id);
  try {
    let question = await API.getQuestion(id);
    return {
      question,
    };
  } catch (e) {
    await goto("/");
    await Swal.fire({
      title: "Error",
      text: "Poll not found",
      icon: "error",
    });
  }
};
