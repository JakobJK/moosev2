<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { Editor } from '@tiptap/core'
  import StarterKit from '@tiptap/starter-kit'
  import Placeholder from '@tiptap/extension-placeholder'

  let {
    value = $bindable(''),
    editor = $bindable(),
    placeholder = 'Write something...'
  } = $props()

  let element = $state<HTMLDivElement>()

  onMount(() => {
    editor = new Editor({
      element: element,
      extensions: [StarterKit, Placeholder.configure({ placeholder })],
      content: value,
      editorProps: {
        attributes: {
          class:
            'prose prose-invert max-w-none focus:outline-none min-h-[200px] p-6 text-slate-300'
        }
      },
      onUpdate: ({ editor }) => {
        value = editor.getHTML()
      },
      onTransaction: () => {
        editor = editor
      }
    })
  })

  onDestroy(() => {
    if (editor) editor.destroy()
  })

  const btnClass =
    'p-2 px-3 rounded text-slate-500 hover:bg-slate-800 hover:text-slate-200 transition-all flex items-center justify-center min-w-[36px]'
  const activeClass = 'text-blue-400 bg-slate-800 shadow-inner'
</script>

<div
  class="bg-slate-900 border border-slate-800 rounded-xl overflow-hidden shadow-2xl focus-within:border-blue-500/50 transition-colors"
>
  <div
    class="bg-slate-950/50 border-b border-slate-800 p-1.5 flex flex-wrap gap-1"
  >
    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleBold().run()}
      class="{btnClass} {editor?.isActive('bold') ? activeClass : ''}"
    >
      <span class="font-bold">B</span>
    </button>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleItalic().run()}
      class="{btnClass} {editor?.isActive('italic') ? activeClass : ''}"
    >
      <span class="italic text-lg">I</span>
    </button>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleStrike().run()}
      class="{btnClass} {editor?.isActive('strike') ? activeClass : ''}"
    >
      <span class="line-through">S</span>
    </button>

    <div class="w-[1px] h-6 bg-slate-800 mx-1 self-center"></div>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleBulletList().run()}
      class="{btnClass} {editor?.isActive('bulletList') ? activeClass : ''}"
    >
      <span class="text-xs">● List</span>
    </button>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleOrderedList().run()}
      class="{btnClass} {editor?.isActive('orderedList') ? activeClass : ''}"
    >
      <span class="text-xs">1. List</span>
    </button>

    <div class="w-[1px] h-6 bg-slate-800 mx-1 self-center"></div>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleCodeBlock().run()}
      class="{btnClass} {editor?.isActive('codeBlock') ? activeClass : ''}"
    >
      <span class="font-mono text-xs">{'<>'}</span>
    </button>

    <button
      type="button"
      onclick={() => editor?.chain().focus().toggleBlockquote().run()}
      class="{btnClass} {editor?.isActive('blockquote') ? activeClass : ''}"
    >
      <span class="font-serif text-lg leading-none">“</span>
    </button>
  </div>

  <div bind:this={element}></div>
</div>

<style>
  :global(.tiptap p.is-editor-empty:first-child::before) {
    color: #475569;
    content: attr(data-placeholder);
    float: left;
    height: 0;
    pointer-events: none;
  }

  :global(.tiptap ul) {
    list-style-type: disc;
    padding-left: 1.5rem;
  }

  :global(.tiptap ol) {
    list-style-type: decimal;
    padding-left: 1.5rem;
  }
</style>
