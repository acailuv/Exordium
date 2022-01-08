const HEADERS = {
  "Content-Type": "application/json",
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, PUT, POST, DELETE, HEAD, OPTIONS",
};

const METHODS = {
  GET: "GET",
  POST: "POST",
  DELETE: "DELETE",
};

const MODE = "cors";

const DEFAULT_OPTIONS = {
  headers: HEADERS,
  mode: MODE,
};

function fetchAndResolve(link: string, opt: RequestInit) {
  return fetch(link, opt)
    .then((response) => response.json())
    .then((data) => {
      if (!data.is_success) {
        throw new Error(data.msg);
      }

      return data;
    })
    .catch((error) => {
      console.log(opt.method, link, error);
    });
}

export function Get(link: string) {
  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.GET,
  } as RequestInit;

  return fetchAndResolve(link, opt);
}

export function Post(link: string, body: object) {
  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.POST,
    body: JSON.stringify(body),
  } as RequestInit;

  return fetchAndResolve(link, opt);
}