<script lang="ts">
    let { data } = $props();

    const BW = 200;
    const BH = 50;
    const VGAP = 100;
    const HGAP = 40;

    let transform = $state({ x: 500, y: 100 }); // Moved start slightly
    let isPanning = $state(false);

    function getNodeWidth(node): number {
        if (!node.children || node.children.length === 0) return BW;
        const childrenWidth = node.children.reduce(
            (acc, child) => acc + getNodeWidth(child),
            0,
        );
        const gaps = (node.children.length - 1) * HGAP;
        return Math.max(BW, childrenWidth + gaps);
    }

    function getChildX(parentX, parentNode, childIndex) {
        const totalChildrenWidth =
            parentNode.children.reduce(
                (acc, child) => acc + getNodeWidth(child),
                0,
            ) +
            (parentNode.children.length - 1) * HGAP;

        // StartX is now the LEFT edge of the entire child row, relative to the parent's CENTER
        let startX = parentX + BW / 2 - totalChildrenWidth / 2;

        let offset = 0;
        for (let i = 0; i < childIndex; i++) {
            offset += getNodeWidth(parentNode.children[i]) + HGAP;
        }

        // Position the child so its center lines up with its allocated slot
        return (
            startX +
            offset +
            getNodeWidth(parentNode.children[childIndex]) / 2 -
            BW / 2
        );
    }
</script>

{#snippet renderNode(node, x, y)}
    <g transform="translate({x}, {y})">
        <rect
            width={BW}
            height={BH}
            rx="6"
            class="fill-slate-900 stroke-slate-700 hover:stroke-blue-500 transition-colors shadow-2xl"
        />
        <text
            x={BW / 2}
            y={BH / 2 + 4}
            text-anchor="middle"
            class="fill-slate-100 text-[11px] font-bold uppercase tracking-wider pointer-events-none"
        >
            {node.label}
        </text>
    </g>

    {#if node.children}
        {#each node.children as child, i}
            {@const childX = getChildX(x, node, i)}
            {@const childY = y + BH + VGAP}

            <path
                d="M {x + BW / 2} {y + BH} 
                   C {x + BW / 2} {y + BH + 40}, 
                     {childX + BW / 2} {childY - 40}, 
                     {childX + BW / 2} {childY}"
                fill="none"
                stroke="#334155"
                stroke-width="2"
            />

            {@render renderNode(child, childX, childY)}
        {/each}
    {/if}
{/snippet}

<div
    class="w-full h-[700px] bg-slate-950 overflow-hidden cursor-grab active:cursor-grabbing border border-slate-800 rounded-xl relative"
    onmousedown={() => (isPanning = true)}
    onmousemove={(e) => {
        if (isPanning) {
            transform.x += e.movementX;
            transform.y += e.movementY;
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
