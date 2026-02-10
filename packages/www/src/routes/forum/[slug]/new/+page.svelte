<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { Editor } from "@tiptap/core";
    import StarterKit from "@tiptap/starter-kit";
    import Placeholder from "@tiptap/extension-placeholder";

    let { data } = $props();

    let element = $state<HTMLDivElement>();
    let editor = $state<Editor>();
    let title = $state("");

    onMount(() => {
        editor = new Editor({
            element: element,
            extensions: [
                StarterKit,
                Placeholder.configure({
                    placeholder:
                        "What are you working on? (Markdown supported...)",
                }),
            ],
            editorProps: {
                attributes: {
                    class: "prose prose-invert max-w-none focus:outline-none min-h-[300px] p-6 text-slate-300",
                },
            },
            onTransaction: () => {
                editor = editor;
            },
        });
    });

    onDestroy(() => {
        if (editor) editor.destroy();
    });

    const formatSlug = (slug: string) => slug.replace(/-/g, " ");
</script>

<div class="max-w-4xl mx-auto p-8">
    <nav
        class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-500 mb-4"
    >
        <a
            href="/forum/{data.slug}"
            class="hover:text-blue-500 transition-colors"
            >Back to {formatSlug(data.slug)}</a
        >
    </nav>

    <h1 class="text-4xl font-black text-white mb-8">Create New Thread</h1>

    <div class="space-y-6">
        <input
            bind:value={title}
            placeholder="Thread Title"
            class="w-full bg-slate-900 border border-slate-800 rounded-xl p-5 text-2xl font-bold text-white focus:outline-none focus:border-blue-500 transition-all shadow-inner"
        />

        <div
            class="bg-slate-900 border border-slate-800 rounded-xl overflow-hidden shadow-2xl"
        >
            <div
                class="bg-slate-950/50 border-b border-slate-800 p-2 flex gap-1"
            >
                <button
                    onclick={() => editor?.chain().focus().toggleBold().run()}
                    class="p-2 px-3 rounded hover:bg-slate-800 transition-colors {editor?.isActive(
                        'bold',
                    )
                        ? 'text-blue-400 bg-slate-800'
                        : 'text-slate-500'}"
                >
                    <b>B</b>
                </button>
                <button
                    onclick={() =>
                        editor?.chain().focus().toggleCodeBlock().run()}
                    class="p-2 px-3 rounded hover:bg-slate-800 transition-colors {editor?.isActive(
                        'codeBlock',
                    )
                        ? 'text-blue-400 bg-slate-800'
                        : 'text-slate-500'}"
                >
                    <span class="font-mono text-sm">{"<>"}</span>
                </button>
            </div>

            <div bind:this={element}></div>
        </div>

        <div class="flex justify-end gap-4">
            <button
                class="bg-blue-600 hover:bg-blue-500 text-white px-10 py-3 rounded-xl font-black shadow-lg shadow-blue-900/40 transition-all active:scale-95"
                onclick={() => console.log(editor?.getJSON())}
            >
                Post Thread
            </button>
        </div>
    </div>
</div>

<style>
    :global(.tiptap p.is-editor-empty:first-child::before) {
        color: #475569;
        content: attr(data-placeholder);
        float: left;
        height: 0;
        pointer-events: none;
    }
</style>
