<script lang="ts">
  import DateDisplay from '$lib/elements/DateDisplay.svelte'
  import BreadCrumbs from '$lib/components/BreadCrumbs.svelte'
  import RichTextEditor from '$lib/components/RichTextEditor.svelte'

  let props = $props()
  let replyContent = $state('')

  function handleReply() {
    if (!replyContent.trim()) return
    // Logic for submission goes here
  }
</script>

<div class="max-w-6xl mx-auto p-6">
  <BreadCrumbs
    forumSlug={props.data.slug}
    forumTitle={props.data.title}
    threadTitle={props.data.threadData?.title}
  />

  <h1 class="text-3xl font-black text-white mb-8">
    {props.data.threadData.title}
  </h1>

  <div class="space-y-6 mb-12">
    {#each props.data.threadData.posts as post}
      <article
        class="grid grid-cols-[200px_1fr] bg-slate-900 border border-slate-800 rounded-xl overflow-hidden shadow-xl"
      >
        <div
          class="bg-slate-950/40 p-6 flex flex-col items-center border-r border-slate-800"
        >
          <div
            class="w-20 h-20 bg-slate-800 rounded-2xl flex items-center justify-center text-3xl font-bold text-blue-500 border border-slate-700 shadow-inner"
          >
            {post.avatar}
          </div>
          <a
            href="/user/{post.author}"
            class="mt-4 font-black text-slate-100 hover:text-blue-400 transition-colors"
          >
            {post.author}
          </a>
          <span
            class="text-[9px] bg-slate-800 text-slate-400 px-2 py-0.5 rounded-full mt-2 uppercase tracking-widest"
          >
            {post.role}
          </span>
        </div>

        <div class="p-8 flex flex-col">
          <div
            class="flex justify-between items-center mb-6 text-[10px] font-mono text-slate-600 uppercase tracking-widest border-b border-slate-800 pb-4"
          >
            <div class="flex gap-1">
              <span>Posted</span>
              <DateDisplay date={post.timestamp} />
            </div>
            <span>Post #{post.id}</span>
          </div>

          <div class="text-slate-300 leading-relaxed text-lg">
            {post.content}
          </div>

          <div class="mt-auto pt-8 flex justify-end gap-6">
            <button
              class="text-[10px] font-black text-slate-500 hover:text-white uppercase tracking-widest"
              >Quote</button
            >
            <button
              class="text-[10px] font-black text-slate-500 hover:text-white uppercase tracking-widest"
              >Report</button
            >
          </div>
        </div>
      </article>
    {/each}
  </div>

  <section class="mt-12 pt-8 border-t border-slate-800 max-w-4xl m-auto">
    <RichTextEditor
      bind:value={replyContent}
      placeholder="Write your response..."
    />
    <div class="mt-6 flex justify-end">
      <button
        onclick={handleReply}
        class="bg-blue-600 hover:bg-blue-500 text-white font-black uppercase text-xs tracking-[0.2em] px-8 py-4 rounded-lg transition-all active:scale-95 shadow-lg shadow-blue-900/20"
      >
        Post Reply
      </button>
    </div>
  </section>
</div>
