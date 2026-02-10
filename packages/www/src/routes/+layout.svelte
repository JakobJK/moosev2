<script>
    import "../app.css";
    import favicon from "$lib/assets/favicon.svg";
    import Nav from "$lib/components/Nav.svelte";
    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<Nav />
<div class="min-h-screen bg-slate-900 text-white">
    {@render children()}
</div>
