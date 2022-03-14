<script lang="ts">
import Router, { link } from "svelte-spa-router";
import { routes } from "./routes.js";

import Header from "./Header.svelte";
import Footer from "./Footer.svelte";
</script>

<main>
	<Header/>	
	<Router {routes}/>
	<Footer/>
</main>

<style>
	:global(html),
	:global(body) {
    	height: 100%;
    	display: flex; 
    	flex-direction: column;
		padding: 0;
	}
</style>