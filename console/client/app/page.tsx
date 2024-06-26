'use client';
import useSWR from "swr";
import React, {useState} from "react";
import useSWRMutation from "swr/mutation";
import {toast, ToastContainer} from "react-toastify";

import 'react-toastify/dist/ReactToastify.css';

const linkFetcher = async (url: string) => {
  const res = await fetch(url);
  if (!res.ok) {
    window.location.href = "/api/auth/login";
    return [] as Link[];
  }

  return await res.json() as Promise<Link[]>;
}
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

async function createLink(url: string, {arg}: { arg: Link }) {
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

async function deleteLink(link: Link): Promise<boolean> {
  const res = await fetch(`/api/links/${link.id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    toast("リンクの削除に失敗しました", {
      type: "error",
      position: "bottom-right"
    });
    return false;
  }

  toast("リンクを削除しました！", {
    type: "success",
    position: "bottom-right"
  });

  return true;
}

async function copyToClipboard(text: string) {
  if (navigator.clipboard) {
    await navigator.clipboard.writeText(text);
    toast("クリップボードにコピーしました！", {
      type: "success",
      position: "bottom-right"
    });
  } else {
    toast("クリップボードへのコピーに失敗しました", {
      type: "error",
      position: "bottom-right"
    });
  }
}

export default function Home() {
  const {data, mutate} = useSWR("/api/links", linkFetcher);
  const {data: metadata} = useSWR("/api/metadata", metadataFetcher);

  const [link, setLink] = useState({} as Link);
  const {trigger: createTrigger, isMutating} = useSWRMutation("/api/links", createLink);

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
            await createTrigger(link);
            await mutate();
          }}>
          {isMutating ? "作成中..." : "リンクを作成する！"}
        </button>
      </div>
      {data?.map((link) => {
        return (
          <div className={"mx-3 md:mx-10 my-5 p-10 rounded-2xl shadow-xl border border-gray-100"} key={link.id}>
            <div className={"flex-row md:flex items-center justify-between"}>
              <div className={"md:w-3/4 w-full flex-row items-center justify-between"}>
                <div className={"flex items-center break-all"}>
                  <span className={"text-2xl w-full mb-3"}>{`${metadata?.link_host}/${link.path}`}</span>
                </div>
                <div className={"flex items-center break-all"}>
                  <span className={"mb-1"}>{link.url}</span>
                </div>
                <div className={"flex items-center break-all"}>
                  <span>{link.description}</span>
                </div>
              </div>
              <div className={"md:w-1/2 mt-5 md:mt-0 w-full flex items-center justify-end"}>
                <button
                  className={"py-1 px-3 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition"}
                  onClick={async () => {
                    await copyToClipboard(`${metadata?.link_host}/${link.path}`);
                  }}>
                  コピー
                </button>
                <button
                  className={"ml-2 py-1 px-3 bg-red-500 text-white rounded-md hover:bg-red-600 transition"}
                  onClick={async _ => {
                    const res = await deleteLink(link);
                    if (res) {
                      await mutate();
                    }
                  }}>
                  削除
                </button>
              </div>
            </div>
          </div>
        );
      })}
      <ToastContainer/>
    </main>
  );
}
