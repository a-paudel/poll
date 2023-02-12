<script lang="ts">
  import { API } from "../../../api";
  import type { IAnswer } from "../../../models";
  import type { PageData } from "./$types";
  import Swal from "sweetalert2";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  export let data: PageData;
  let question = data.question;

  async function vote(answer: IAnswer) {
    if (localStorage.getItem(question.id.toString()) === null) {
      question = await API.voteQuestion(question.id, answer.id);
      // add to local storage
      localStorage.setItem(question.id.toString(), answer.id.toString());
      //   goto view page
      await goto(`/${question.id}/view`);
      //   await invalidateAll();
      await Swal.fire({
        icon: "success",
        title: "Success!",
        text: "Your vote has been recorded!",
      });
    } else {
      // show error
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "You have already voted for this answer!",
      });
    }
  }

  async function init() {
    // check if user has voted
    if (localStorage.getItem(question.id.toString()) !== null) {
      // user has voted, goto results page
      await goto(`/${question.id}/view`);
      await Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "You have already voted for this answer!",
      });
    }
  }

  onMount(() => {
    init();
  });
</script>

<div
  class="h-screen bg-dark text-light flex flex-col gap4 items-center justify-center px-4"
>
  <h1 class="text-center pt10">{question.question}</h1>

  <div class="flex flex-col text-center gap4 w-full">
    {#each question.answers as answer}
      <!-- content here -->
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <div
        class="bg-dark-3 px2 py10 rounded cursor-pointer relative top-0 hover:-top-1 transition-all"
        on:click={() => vote(answer)}
      >
        <h2>{answer.answer}</h2>
      </div>
    {/each}
  </div>
</div>
