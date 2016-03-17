package main;

import "fmt";

p := make([]int, 1000001);
seen := make(bool, 1000001);

func find_set(x int) int {
  return (p[x] == x)? x : p[x] = find_set(p[x]);
}

func union_set(x int, y int) {
  px := find_set(x);
  py := find_set(y);
  if px != py {
    p[px] = py;
  }
}

func init(n int) {
  for i := 0; i <= n; i++ {
    p[i] = i;
  }
}

func dfs(u int, g[][]int, vi bool) {
  vi[u] = true;
  fmt.Println(u);
  
  for i := 0; i < len(g[u]); i++ {
    if (!vi[g[u][i]]) {
      dfs(g[u][i], g, vi);
    }
  }
}

func main() {
  var n, u, v int;
  fmt.Scanf("%d", &n);
  //g := make([][]int, n);
  //vi := make([]bool, n);
  init(n);
  
  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d", &u, &v);
    union_set(u, v);
    //g[u] = append(g[u], v);
    //g[v] = append(g[v], u);
  }
  
  fmt.Println("Parents:");
  for i := 0; i < n; i++ {
    if !seen[find_set(i)] {
      seen[find_set(i)] = true;
      fmt.Printf("%d", i);
    }
  }
  fmt.Println();
  
  //fmt.Scanf("%d", &init);
  //dfs(init, g, vi);
}