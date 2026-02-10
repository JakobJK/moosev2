<script lang="ts">
    // Svelte 5 props
    let { forumSlug, forumTitle, threadTitle = "" } = $props();

    $effect(() => {
        if (!forumSlug || !forumTitle) {
            print("Breadcrumbs warning: forumSlug or forumTitle is missing.");
        }
    });
</script>

<nav class="flex text-sm items-center gap-2 text-slate-500 mb-6" aria-label="Breadcrumb">
    <a href="/forum" class="hover:text-blue-500 transition-colors">Forums</a>
    
    <span>/</span>
    
    {#if !threadTitle}
        <span class="text-slate-300">{forumTitle}</span>
    {:else}
        <a href="/forum/{forumSlug}" class="hover:text-blue-500 transition-colors">
            {forumTitle}
        </a>
        <span>/</span>
        <span class="text-slate-300 truncate">{threadTitle}</span>
    {/if}
</nav>
