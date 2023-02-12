<script lang="ts">
  import { onMount } from "svelte";
  import { API } from "../../../api";
  import type { IQuestion } from "../../../models";
  import type { PageData } from "./$types";
  import { Chart } from "chart.js/auto";
  import { browser } from "$app/environment";
  import Swal from "sweetalert2";

  export let data: PageData;
  let question: IQuestion = data.question;
  let graph: Chart;
  $: {
    question;
    if (graph) {
      updateGraph();
    }
  }
  let voteLink = "";
  $: {
    if (browser) {
      voteLink = `${window.location.origin}/${question.id}/vote`;
    }
  }
  function init() {
    let ws = API.listenQuestion(question.id);
    ws.onmessage = (e) => {
      question = JSON.parse(e.data);
    };
  }

  function updateGraph() {
    graph.data.labels = question.answers.map((a) => a.answer);
    graph.data.datasets[0].data = question.answers.map((a) => a.votes);
    graph.update();
  }

  function initGraph() {
    let graphElement = document.getElementById("graph") as HTMLCanvasElement;
    graph = new Chart(graphElement, {
      type: "bar",
      options: {
        plugins: { legend: { display: false } },
      },
      data: {
        labels: question.answers.map((a) => a.answer),
        datasets: [
          {
            data: question.answers.map((a) => a.votes),
            indexAxis: "y",
          },
        ],
      },
    });
  }

  function copyHandler() {
    try {
      navigator.clipboard.writeText(voteLink);
      Swal.fire({
        icon: "success",
        title: "Copied!",
        text: "The link has been copied to your clipboard",
        timer: 1000,
        timerProgressBar: true,
        showConfirmButton: false,
        position: "bottom",
      });
    } catch (e) {
      console.error(e);
    }
  }

  onMount(() => {
    init();
    initGraph();
  });
</script>

<div
  class="h-screen bg-dark text-light flex flex-col items-center justify-center px-4"
>
  <h1 class="text-center pt10">{question.question}</h1>
  <canvas id="graph" class="w-5/6" />
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div class="mt10">Vote at</div>
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <code
    class="mt5 text-lg bg-dark-1 p3 rounded cursor-pointer select-all"
    on:click={copyHandler}
  >
    {voteLink}
  </code>
</div>
