<script lang="ts">
    let props = $props();
</script>

<div class="max-w-6xl mx-auto p-6">
    <h1 class="text-3xl font-bold mb-8">Community Forum</h1>

    {#each props.data.categories as category}
        <section class="mb-10">
            <h2
                class="text-xl font-semibold text-blue-400 mb-4 border-b border-slate-800 pb-2"
            >
                {category.name}
            </h2>

            <div class="grid gap-4">
                {#each category.forums as forum}
                    <div
                        class="bg-slate-800 p-4 rounded-lg border border-slate-700 flex justify-between items-center hover:border-slate-500 transition-colors"
                    >
                        <div class="flex-1 min-w-0">
                            <h3
                                class="font-bold text-lg hover:underline cursor-pointer text-slate-100"
                            >
                                <a href="/forum/{forum.slug}">{forum.title}</a>
                            </h3>
                            <p class="text-sm text-slate-400 mb-2">
                                {forum.description}
                            </p>
                            {#if forum.latestPost}
                                <p class="text-xs text-slate-500 truncate">
                                    Latest: <span class="text-slate-300"
                                        >"{forum.latestPost.title}"</span
                                    >
                                    by
                                    <a
                                        href="user/{forum.latestPost.username}"
                                        class="text-blue-400"
                                        >{forum.latestPost.username}</a
                                    >
                                </p>
                            {:else}
                                <p class="text-xs text-slate-600 italic">
                                    No posts yet
                                </p>
                            {/if}
                        </div>

                        <div class="text-right ml-6 shrink-0">
                            <span
                                class="block font-mono text-xl text-blue-500 leading-none"
                                >{forum.postCount}</span
                            >
                            <span
                                class="text-[10px] uppercase tracking-wider text-slate-500 font-bold"
                                >Posts</span
                            >
                        </div>
                    </div>
                {/each}
            </div>
        </section>
    {/each}
</div>
