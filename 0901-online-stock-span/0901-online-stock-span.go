type Stock struct {
    Price int
    Span int
}

type StockSpanner struct {
    stack []Stock
}


func Constructor() StockSpanner {
    return StockSpanner{
        stack: []Stock{},
    }
}


func (this *StockSpanner) Next(price int) int {
    span := 1

    for len(this.stack) > 0 && this.stack[len(this.stack) - 1].Price <= price {
        span += this.stack[len(this.stack) - 1].Span
        this.stack = this.stack[:len(this.stack) - 1]
    }

    this. stack = append(this.stack, Stock{Price: price, Span: span})
    
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */