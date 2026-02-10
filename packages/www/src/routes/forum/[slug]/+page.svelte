<script lang="ts">

    import BreadCrumbs from '$lib/components/BreadCrumbs.svelte';
    let { data } = $props();

</script>

<div class="max-w-6xl mx-auto p-6">
    
<BreadCrumbs 
    forumSlug={data.slug}
    forumTitle={data.forumName}
/>

    <div class="flex justify-between items-end mb-8">
        <div>
            <h1 class="text-4xl font-black text-white">{data.forumName}</h1>
            <p class="text-slate-400 mt-2">{data.description}</p>
        </div>
        <a href="/forum/{data.slug}/new" class="bg-blue-600 hover:bg-blue-500 text-white px-6 py-2 rounded font-bold transition-all">
            New Thread
        </a>
    </div>

    <div class="bg-slate-900 border border-slate-800 rounded-lg overflow-hidden">
        <div class="grid grid-cols-12 bg-slate-950 p-4 text-[10px] uppercase tracking-widest font-bold text-slate-500 border-b border-slate-800">
            <div class="col-span-7 lg:col-span-8">Topic</div>
            <div class="hidden lg:block col-span-2 text-center">Stats</div>
            <div class="col-span-5 lg:col-span-2 text-right lg:text-left">Last Post</div>
        </div>

        {#each data.threads as thread}
            <div class="grid grid-cols-12 p-4 border-b border-slate-800/50 hover:bg-slate-800/30 transition-colors items-center">
                <div class="col-span-7 lg:col-span-8">
                    <div class="flex items-center gap-2">
                        {#if thread.isPinned}
                            <span class="text-[10px] bg-blue-900/50 text-blue-400 px-1.5 py-0.5 rounded border border-blue-800/50">PINNED</span>
                        {/if}
                        <a href="/forum/{data.slug}/{thread.id}" class="text-lg font-bold text-slate-200 hover:text-blue-400 transition-colors">
                            {thread.title}
                        </a>
                    </div>
                    <p class="text-xs text-slate-500 mt-1">
                        by <a href="/user/{thread.author}" class="text-slate-400 hover:text-blue-400">{thread.author}</a>
                    </p>
                </div>

                <div class="hidden lg:flex col-span-2 justify-center gap-6 text-center">
                    <div>
                        <span class="block text-slate-300 font-mono text-sm">{thread.replies}</span>
                        <span class="text-[9px] uppercase text-slate-600">Replies</span>
                    </div>
                    <div>
                        <span class="block text-slate-300 font-mono text-sm">{thread.views}</span>
                        <span class="text-[9px] uppercase text-slate-600">Views</span>
                    </div>
                </div>

                <div class="col-span-5 lg:col-span-2 text-right lg:text-left">
                    <span class="block text-sm text-slate-300">{thread.lastPost}</span>
                    <span class="text-[10px] text-slate-600">by Anonymous</span>
                </div>
            </div>
        {/each}
    </div>
</div>
