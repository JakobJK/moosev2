<script lang="ts">
    import DateDisplay from '$lib/elements/DateDisplay.svelte';
    let { data } = $props();
</script>

<div class="max-w-6xl mx-auto p-8">
    <div class="bg-slate-900 border border-slate-700 rounded-lg overflow-hidden mb-8">
        <div class="h-32 bg-gradient-to-r from-blue-900 to-slate-800"></div>
        <div class="px-8 pb-8">
            <div class="relative flex justify-between items-end -mt-12 mb-6">
                <div class="w-32 h-32 bg-slate-800 border-4 border-slate-900 rounded-lg flex items-center justify-center text-5xl font-bold text-blue-500 shadow-2xl">
                    {data.username[0]}
                </div>
                <div class="flex gap-3">
                    <button class="bg-blue-600 hover:bg-blue-500 text-white px-6 py-2 rounded font-bold text-sm transition-all">Follow</button>
                    <button class="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded font-bold text-sm transition-all">Message</button>
                </div>
            </div>
            
            <h1 class="text-4xl font-black text-white tracking-tight">{data.username}</h1>
            <p class="text-slate-400 mt-2 max-w-2xl">{data.bio}</p>
        </div>
    </div>

    <div class="grid grid-cols-12 gap-8">
        <div class="col-span-4 space-y-4">
            <div class="bg-slate-900 border border-slate-700 p-6 rounded-lg">
                <h3 class="text-xs font-bold text-slate-500 uppercase tracking-widest mb-4">User Statistics</h3>
                <div class="grid grid-cols-2 gap-4">
                    <div>
                        <p class="text-2xl font-mono text-white leading-none">{data.stats.posts}</p>
                        <p class="text-[10px] text-slate-500 uppercase mt-1">Posts</p>
                    </div>
                    <div>
                        <p class="text-2xl font-mono text-white leading-none">{data.stats.reputation}</p>
                        <p class="text-[10px] text-slate-500 uppercase mt-1">Reputation</p>
                    </div>
                </div>
                <div class="mt-6 pt-6 border-t border-slate-800 text-xs text-slate-500">
                    Joined: <DateDisplay date={data.joinedDate} className="text-slate-300 ml-1" />
                </div>
            </div>
        </div>

        <div class="col-span-8">
            <h3 class="text-xl font-bold text-white mb-6">Recent Activity</h3>
            <div class="border-l-2 border-slate-800 ml-4 space-y-8">
                {#each data.recentActivity as item}
                    <div class="relative pl-8">
                        <div class="absolute w-4 h-4 bg-blue-600 rounded-full -left-[9px] top-1 border-4 border-slate-900"></div>
                        <div class="bg-slate-900 border border-slate-700 p-5 rounded-lg hover:border-blue-500 transition-colors cursor-pointer">
                            <div class="flex justify-between items-start mb-1">
                                <h4 class="font-bold text-slate-100">{item.title}</h4>
                                <DateDisplay date={item.date} className="text-[10px] font-mono text-slate-500 uppercase" />
                            </div>
                            <p class="text-sm text-slate-400">
                                User posted a new <span class="text-blue-400">{item.type}</span> in <span class="underline">Hard Surface Modeling</span>
                            </p>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</div>
