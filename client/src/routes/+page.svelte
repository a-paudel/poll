<script lang="ts">
  import { goto } from "$app/navigation";
  import { API } from "../api";

  let question: string = "";
  let answers: string[] = [];

  function addAnswerHandler() {
    answers = [...answers, ""];
  }
  function removeAnswerHandler(index: number) {
    answers = answers.filter((_, i) => i !== index);
  }

  async function submitHandler() {
    // check if question is empty
    if (question === "") {
      alert("Please add a question");
      return;
    }
    // remove empty answers
    answers = answers.filter((answer) => answer !== "");
    // check if answers are empty
    if (answers.length < 2) {
      alert("Please add at least two answers");
      return;
    }

    // create poll
    let data = await API.createQuestion({ question, answers });
    // redirect to poll view page
    goto(`/${data.id}/view`);
  }
</script>

<h1 class="fixed w-screen text-light text-center text-6xl top-5">Poll</h1>

<div class="flex flex-col md:flex-row">
  <!-- question -->
  <div
    class="bg-dark-9 min-h-screen w-full md:w-1/2 flex items-center justify-center"
  >
    <input
      type="text"
      placeholder="Question"
      class="p2 text-4xl text-light bg-transparent border-0 text-center focus:outline-none"
      autofocus
      bind:value={question}
    />
  </div>

  <div
    class="bg-dark-7 min-h-screen w-full md:w-1/2 flex flex-col items-center justify-center"
  >
    <!-- answers -->
    {#each answers as answer, index}
      <div class="flex items-center justify-evenly gap2 w-5/6">
        <span class="text-light font-bold">{index + 1}</span>
        <input
          type="text"
          bind:value={answer}
          class="w-5/6 p2 text-2xl rounded bg-dark-9 border-0  text-light focus:outline-none"
          autofocus
        />
        <button
          class="bg-dark-9 hover:bg-dark-8 text-lg text-light font-bold hfull aspect-square rounded border-0 cursor-pointer"
          on:click={() => removeAnswerHandler(index)}
        >
          x
        </button>
      </div>
      <div class="h5" />
    {/each}
    <button
      class="text-center p2 text-lg font-bold bg-dark-9 hover:bg-dark-8 border-0 rounded text-light cursor-pointer"
      on:click={addAnswerHandler}
    >
      Add Answer
    </button>
  </div>
</div>

<div class="fixed bottom-0 w-screen flex justify-center">
  <button
    class="text-xl font-bold p2 px-4 border-0 rounded-t bg-dark hover:bg-dark-8 active:bg-dark-9  text-light cursor-pointer"
    on:click={submitHandler}
  >
    Create Poll
  </button>
</div>
