import Header from "../components/Header"
import TicketRow from "../components/TicketRow"

function TicketList() {
  const items = [
    {
      id: 1,
      title: 'タイトル',
      tags: ['status:open', 'assignee:yosiopp', 'type:bug', 'document']
    },
    {
      id: 2,
      title: 'タイトルタイトルタイトル',
      tags: []
    }
  ]

  return (
    <>
      <Header/>
      <div>絞り込み条件</div>
      <div className="flex text-gray-500 border-b">
        <div className="flex-none w-16 py-1 pl-4">id</div>
        <div className="w-2/4 py-1">title</div>
        <div className="flex-1 py-1 pr-4">tags</div>
      </div>
      { items.map((item) => <TicketRow key={item.id} data={item} onClick={(id) => { console.log(id)}}/>)}
    </>
  )
}

export default TicketList