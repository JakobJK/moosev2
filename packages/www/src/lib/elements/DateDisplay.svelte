<script lang="ts">
    interface Props {
        date: string | Date;
        className?: string;
    }

    let { date, className = "" }: Props = $props();

    const formattedDate = $derived.by(() => {
        const d = new Date(date);
        if (isNaN(d.getTime())) return "Invalid Date";

        return new Intl.DateTimeFormat('en-US', {
            month: 'short',
            day: '2-digit',
            year: 'numeric'
        }).format(d);
    });
</script>

<span class={className}>
    {formattedDate}
</span>
