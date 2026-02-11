<script lang="ts">
  import Chapter from '$lib/components/Chapter.svelte'
  let { data } = $props()

  const BW = 200 // Minimum Width
  const VGAP = 100
  const HGAP = 40

  let transform = $state({ x: 500, y: 100 })
  let isPanning = $state(false)

  // 1. We still need the math to know where to place the ForeignObjects
  function getNodeWidth(node): number {
    if (!node.children || node.children.length === 0) return BW
    const childrenWidth = node.children.reduce(
      (acc, child) => acc + getNodeWidth(child),
      0
    )
    const gaps = (node.children.length - 1) * HGAP
    return Math.max(BW, childrenWidth + gaps)
  }

  function getChildX(parentX, parentNode, childIndex) {
    const totalChildrenWidth =
      parentNode.children.reduce((acc, child) => acc + getNodeWidth(child), 0) +
      (parentNode.children.length - 1) * HGAP

    let startX = parentX + BW / 2 - totalChildrenWidth / 2
    let offset = 0
    for (let i = 0; i < childIndex; i++) {
      offset += getNodeWidth(parentNode.children[i]) + HGAP
    }

    return (
      startX +
      offset +
      getNodeWidth(parentNode.children[childIndex]) / 2 -
      BW / 2
    )
  }
</script>

{#snippet renderNode(node, x, y)}
  {#if node.children}
    {#each node.children as child, i}
      {@const childX = getChildX(x, node, i)}
      {@const childY = y + VGAP}
      <path
        d="M {x + BW / 2} {y} 
           C {x + BW / 2} {y + VGAP / 2}, 
             {childX + BW / 2} {childY - VGAP / 2}, 
             {childX + BW / 2} {childY}"
        fill="none"
        stroke="#334155"
        stroke-width="2"
      />
      {@render renderNode(child, childX, childY)}
    {/each}
  {/if}

  <foreignObject {x} {y} width={BW} height="500" class="overflow-visible">
    <Chapter {node} />
  </foreignObject>
{/snippet}

<div
  class="w-full h-screen bg-slate-950 overflow-hidden cursor-grab active:cursor-grabbing"
  onmousedown={() => (isPanning = true)}
  onmousemove={(e) => {
    if (isPanning) {
      transform.x += e.movementX
      transform.y += e.movementY
    }
  }}
  onmouseup={() => (isPanning = false)}
  onmouseleave={() => (isPanning = false)}
  role="presentation"
>
  <svg class="w-full h-full">
    <g transform="translate({transform.x}, {transform.y})">
      {@render renderNode(data.tree, 0, 0)}
    </g>
  </svg>
</div>
