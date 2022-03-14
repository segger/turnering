<script lang="ts">
import { onMount } from "svelte";

import { link } from 'svelte-spa-router'

let contestName = "";

onMount(async () => {
    let contestId = Cookies.get('registerId');
    
    fetch("/api/contests/"+contestId)
    .then(response => response.json())
    .then(data => {
        if (data) {
            contestName = data.name;
        }
    }).catch(error => {
        console.error(error);
        let results = undefined;
    });

    fetch("/api/results/"+contestId)
    .then(response => response.json())
    .then(data => handleResult(data))
    .catch(error => {
        console.log(error);
    });
});

function handleResult(data) {
    if (data) {
        console.log(data);
    }
}

</script>

<main>
    <div class="container">
        <div class="contestHeader">
            <a href="/" use:link class="home">
                <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-house" viewBox="0 0 16 16">
                    <path fill-rule="evenodd" d="M2 13.5V7h1v6.5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V7h1v6.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 13.5zm11-11V6l-2-2V2.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5z"/>
                    <path fill-rule="evenodd" d="M7.293 1.5a1 1 0 0 1 1.414 0l6.647 6.646a.5.5 0 0 1-.708.708L8 2.207 1.354 8.854a.5.5 0 1 1-.708-.708L7.293 1.5z"/>
                </svg>
            </a>
        </div>
        <h2 id="contestName">{contestName}</h2>
    </div>

    <div id="result">
        
    </div>
</main>

<style>
    .contestHeader {
        display: flex;
    }
    a.home {
        color: black;
        text-decoration: none;
        padding-right: 10px;
        padding-top: 5px;
    }
    .container {
		display: flex;
	}
</style>