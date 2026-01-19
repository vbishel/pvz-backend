import { queryOptions } from "@tanstack/react-query";
import { getMe } from "./getMe";

export const queries = {
    all: () => ["session"],
    me: () => queryOptions({
        queryKey: [...queries.all(), "me"],
        queryFn: () => getMe(),
    })
}
