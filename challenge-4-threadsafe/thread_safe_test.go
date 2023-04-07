package threadsafe_test

//func TestThreadSafety(t *testing.T) {
//	t.Run("SimpleKeyValue", func(t *testing.T) {
//		t.Run("not thread safe", func(t *testing.T) {
//			//mostly panics and passes the test this way, but sometimes passes with a non empty map after the checker
//			simpleKeyValues := simple.NewSliceKeyValues[interfaces.MyCustomInt, interfaces.MyCustomString]()
//
//			got := threadsafe.ThreadSafetyChecker(simpleKeyValues)
//
//			for i := 0; i < 1000; i++ {
//				_, isGot := got.Get(i)
//				if isGot {
//					return
//				}
//			}
//			assert.Fail(t, "should not get here")
//		})
//		t.Run("is thread safe", func(t *testing.T) {
//			simpleKeyValues := simple.NewSliceKeyValues[interfaces.MyCustomInt, interfaces.MyCustomString]()
//
//			got := threadsafe.ThreadSafetyChecker(simpleKeyValues)
//
//			for i := 0; i < 1000; i++ {
//				getValue := got.Get(i)
//				assert.Empty(t, getValue)
//			}
//		})
//	})
//	t.Run("Hash", func(t *testing.T) {
//		t.Run("not thread safe", func(t *testing.T) {
//			//only sometimes is it not threadsafe - ie the test passes
//			hashmap := hashing_solution.NewHashMap()
//
//			got := threadSafetyChecker(hashmap)
//
//			for i := 0; i < 1000; i++ {
//				_, isGot := got.Get(i)
//				if isGot {
//					return
//				}
//			}
//			assert.Fail(t, "should not get here")
//		})
//		t.Run("is thread safe", func(t *testing.T) {
//			hashmap := hashing_solution.NewHashMap()
//			wrapper := NewThreadSafeMap(hashmap)
//
//			got := threadSafetyChecker(wrapper)
//
//			for i := 0; i < 1000; i++ {
//				getValue, _ := got.Get(i)
//				assert.Empty(t, getValue)
//			}
//		})
//	})
//}
