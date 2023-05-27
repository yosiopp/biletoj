type Props = {
  data: {
    id: string
    text: string
    color: string
  };
};

function TagItem({ data }: Props) {
  const isGroup = data.text.indexOf(':') > 0;
  const [parent, child] = isGroup ? data.text.split(':', 2) : [];

  return isGroup ? (
    <span className="rounded-lg bg-neutral-100 py-1 px-2 mr-1">
      <span className="border-r pr-1 text-sm opacity-70">{parent}</span>
      <span className="pl-2">{child}</span>
    </span>
  ) : (
    <span className="rounded-lg bg-neutral-100 py-1 px-2 mr-1">
      {data.text}
    </span>
  );
}

export default TagItem;
