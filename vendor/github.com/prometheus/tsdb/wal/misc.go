package wal

// func BenchmarkSyscall(b *testing.B) {
// 	f, err := os.Create("testfile")
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	b.ResetTimer()
// 	b.SetBytes(2333)

// 	buf := make([]byte, 2333)

// 	for i := 0; i < b.N; i++ {
// 		_, err := f.Write(buf)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }

// func TestA(t *testing.T) {
// 	f, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err)
// 	}

// 	go func() {
// 		var buf [pageSize]byte
// 		for i := 0; ; i++ {
// 			binary.BigEndian.PutUint64(buf[:], uint64(i))
// 			if _, err := f.Write(buf[:]); err != nil {
// 				panic(err)
// 			}
// 			fmt.Println("write", i)
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	f2, err := os.Open("testfile")
// 	if err != nil {
// 		panic(err)
// 	}
// 	r := bufio.NewReaderSize(f2, pageSize)
// 	for {
// 		var buf [pageSize]byte
// 		m := 0
// 		for m < pagesPerSegment {
// 			n, err := r.Read(buf[0:])
// 			if err == io.EOF {
// 				time.Sleep(1)
// 			} else if err != nil {
// 				panic(err)
// 			}
// 			m += n
// 			if n > 0 {
// 				fmt.Println("N", n, err)
// 				fmt.Println("m", m)
// 			}
// 		}
// 		x := binary.BigEndian.Uint64(buf[:])
// 		fmt.Println("read", x, fmt.Sprintf("%x", buf)[:32])
// 	}
// }
