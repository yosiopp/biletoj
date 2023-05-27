import TagItem from "./TagItem";

type Props = {
  data: {
    id: number;
    title: string;
    tags: string[];
  };
  onClick: (id: number) => void;
};

function TicketRow({ data, onClick }: Props) {
  const { id, title, tags } = data;
  return (
    <div className="flex hover:bg-slate-100 cursor-pointer" onClick={() => onClick(id)}>
      <div className="flex-none w-16 py-1 pl-4">{id}</div>
      <div className="w-2/4 py-1">{title}</div>
      <div className="flex-1 py-1 pr-4">
        {tags?.map((tag, i) => (
            <TagItem key={i} data={{id: `${i}`, text: tag, color: ''}} />
        ))}
      </div>
    </div>
  );
}

export default TicketRow;
