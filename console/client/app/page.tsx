async function getLinks() {
  // const response = await fetch("/api/links");
  // const data: Link[] = await response.json();
  return [
    {id: "1", url: "https://google.com", shortUrl: "https://beal.ink/foobar", description: "Google"},
    {id: "2", url: "https://twitter.com", shortUrl: "https://beal.ink/foobar", description: "Twitter"},
    {id: "3", url: "https://facebook.com", shortUrl: "https://beal.ink/foobar", description: "Facebook"},
    {id: "4", url: "https://instagram.com", shortUrl: "https://beal.ink/foobar", description: "Instagram"},
    {id: "5", url: "https://youtube.com", shortUrl: "https://beal.ink/foobar", description: "YouTube"},
  ];
}

type Link = {
  id: string;
  url: string;
  shortUrl: string;
  description: string;
}

export default async function Home() {
  const data = await getLinks();

  return (
    <main className={"light"}>
      <div className={"m-10 p-10 rounded-2xl shadow-xl sticky top-0 border border-gray-100 bg-white"}>
        <h1 className={"text-2xl mb-3"}>新しくリンクを作成</h1>
        <input
          type={"text"}
          className={"w-full p-2 border border-gray-200 focus:border-gray-400 focus:outline-none rounded"}
          placeholder={"リンクのURLを入力"}/>
        <button
          className={"mt-3 w-full py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"}>
          リンクを作成する！
        </button>
      </div>
      {data.map((link) => {
        return (
          <div className={"mx-10 my-5 p-10 rounded-2xl shadow-xl border border-gray-100"} key={link.id}>
            <div className={"flex items-center justify-between"}>
              <div className={"w-3/4 flex-row items-center justify-between"}>
                <div className={"flex items-center"}>
                  <span className={"ml-3 text-2xl w-full mb-3"}>{link.shortUrl}</span>
                </div>
                <div className={"flex items-center"}>
                  <span className={"ml-3 mb-1"}>{link.url}</span>
                </div>
                <div className={"flex items-center"}>
                  <span className={"ml-3"}>{link.description}</span>
                </div>
              </div>
              <div className={"w-1/2 flex items-center justify-end"}>
                <button
                  className={"py-1 px-3 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition"}>
                  コピー
                </button>
                <button
                  className={"ml-2 py-1 px-3 bg-red-500 text-white rounded-md hover:bg-red-600 transition"}>
                  削除
                </button>
              </div>
            </div>
          </div>
        );
      })}
    </main>
  );
}
