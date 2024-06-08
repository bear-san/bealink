'use client';
import useSWR from "swr";
import React, {useState} from "react";
import useSWRMutation from "swr/mutation";
import {toast, ToastContainer} from "react-toastify";

import 'react-toastify/dist/ReactToastify.css';

const linkFetcher = (url: string) => fetch(url).then((res) => res.json() as Promise<Link[]>);
const metadataFetcher = (url: string) => fetch(url).then((res) => res.json() as Promise<Metadata>);

type Link = {
  id: string;
  url: string;
  path: string;
  description: string;
}

type Metadata = {
  link_host: string;
}

async function createLink(url: string, { arg }: { arg: Link }) {
  const res = await fetch(url, {
    method: "POST",
    body: JSON.stringify(arg),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    toast("リンクの作成に失敗しました", {
      type: "error",
      position: "bottom-right"
    });
    return res.json();
  }

  toast("リンクを作成しました！", {
    type: "success",
    position: "bottom-right"
  });

  return res.json();
}

export default function Home() {
  const {data, error, isLoading} = useSWR("/api/links", linkFetcher);
  const {data: metadata} = useSWR("/api/metadata", metadataFetcher);

  const [link, setLink] = useState({} as Link);
  const {trigger, isMutating} = useSWRMutation("/api/links", createLink);

  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    switch (e.target.formTarget) {
      case "link":
        link.url = e.target.value;
        break;
      case "path":
        link.path = e.target.value;
        break;
      case "description":
        link.description = e.target.value;
        break;
    }
    setLink(link);
  }

  return (
    <main className={"light"}>
      <div className={"m-3 md:m-10 p-10 rounded-2xl shadow-xl sticky top-0 border border-gray-100 bg-white"}>
        <h1 className={"text-2xl mb-3"}>新しくリンクを作成</h1>
        <input
          type={"text"}
          className={"w-full p-2 border border-gray-200 focus:border-gray-400 focus:outline-none rounded mb-3"}
          placeholder={"リンクのURLを入力"}
          onChange={onChange}
          formTarget={"link"}/>
        <input
          type={"text"}
          className={"w-full p-2 border border-gray-200 focus:border-gray-400 focus:outline-none rounded mb-3"}
          placeholder={"カスタムパス"}
          onChange={onChange}
          formTarget={"path"}/>
        <input
          type={"text"}
          className={"w-full p-2 border border-gray-200 focus:border-gray-400 focus:outline-none rounded"}
          placeholder={"メモ"}
          onChange={onChange}
          formTarget={"description"}/>
        <button
          className={`mt-3 w-full py-2 text-white rounded transition ${isMutating ? "bg-gray-400" : "bg-blue-500 hover:bg-blue-600"}`}
          onClick={async () => {
            console.log(link);
            await trigger(link);
          }}>
          {isMutating ? "作成中..." : "リンクを作成する！"}
        </button>
      </div>
      {data?.map((link) => {
        return (
          <div className={"mx-3 md:mx-10 my-5 p-10 rounded-2xl shadow-xl border border-gray-100"} key={link.id}>
            <div className={"flex-row md:flex items-center justify-between"}>
            <div className={"w-3/4 flex-row items-center justify-between"}>
                <div className={"flex items-center"}>
                  <span className={"text-2xl w-full mb-3"}>{`${metadata?.link_host}/${link.path}`}</span>
                </div>
                <div className={"flex items-center"}>
                  <span className={"mb-1"}>{link.url}</span>
                </div>
                <div className={"flex items-center"}>
                  <span>{link.description}</span>
                </div>
              </div>
              <div className={"md:w-1/2 w-full flex items-center justify-end"}>
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
      <ToastContainer />
    </main>
  );
}
