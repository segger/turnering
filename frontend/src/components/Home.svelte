<script lang="ts">
import { onMount } from "svelte";
    
import { Button } from "sveltestrap";

import { push } from 'svelte-spa-router';
    
let contests = [];
    
onMount(async () => {
    fetch("/api/contests")
    .then(response => response.json())
    .then(data => {
        if (data && data.length) {
            console.log(data);
            contests = data;
        }
    }).catch(error => {
        console.error(error);
    });
});
    
function onClick(href: String, contestId: String) {
    Cookies.set('registerId', contestId);
    push('/result');
}
</script>
    
<main>
    <div class="container">
        <div>
            <h3>Registrera resultat</h3>
            <div id="contests" class="form-group">
                {#if contests !== undefined && contests.length > 0}
                    {#each contests as contest}
                        <Button on:click={() => onClick('register', contest.id)} style="margin: 10px;" color="primary" disabled={contest.enabled}>{contest.name}</Button>
                    {/each}
                {:else}
                    <div>Det finns inga tävlingar</div>
                {/if}
            </div>
            <h3>Visa registrerade</h3>
            <div id="registered" class="form-group">
                {#if contests !== undefined && contests.length > 0}
                    {#each contests as contest}
                        <Button on:click={() => onClick('registered', contest.id)} style="margin: 10px;" color="primary" disabled={contest.enabled}>{contest.name}</Button>
                    {/each}
                {:else}
                    <div>Det finns inga tävlingar</div>
                {/if}
            </div>
            <h3>Visa resultat</h3>
            <div id="result" class="form-group">
                {#if contests !== undefined && contests.length > 0}
                    {#each contests as contest}
                        <Button on:click={() => onClick('result', contest.id)} style="margin: 10px;" color="primary" disabled={!contest.enabled}>{contest.name}</Button>
                    {/each}
                {:else}
                    <div>Det finns inga tävlingar</div>
                {/if}
            </div>
        </div>
    </div>
</main>
    
<style>
</style>