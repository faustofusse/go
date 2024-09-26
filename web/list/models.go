package list

type Row interface {
    GetID() string
}

type SearchOptions struct {
    BasePath string
    Query string
    QueryParsed string
    SortField string
    SortPath string
    Filter string
    Ascending bool
    Limit int64
    Offset int64
}

func (options *SearchOptions) Path() string {
    query := options.Query
    filter := options.Filter
    sort := options.SortPath
    total := len(query) + len(filter) + len(sort)
    path := options.BasePath
    if total == 0 { return path }
    path += "?" 
    // query
    if len(query) > 0 {
        path += "q=" + query
        total -= len(query)
        if total > 0 {
            path += "&" 
        } else {
            return path
        }
    }
    // sort
    if len(sort) > 0 {
        path += "sort=" + sort
        total -= len(sort)
        if total > 0 {
            path += "&" 
        } else {
            return path
        }
    }
    // filter
    if len(filter) > 0 {
        path += "filter=" + filter
        total -= len(filter)
        if total > 0 {
            path += "&" 
        } else {
            return path
        }
    }
    return path
}
