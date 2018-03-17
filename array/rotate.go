package array

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}

	pre := make([]int, 0, len(nums)-k)
	post := make([]int, 0, k)
	for i := 0; i < len(nums)-k; i++ {
		pre = append(pre, nums[i]) // 存储前半部分
	}
	for i := len(nums) - k; i < len(nums); i++ {
		post = append(post, nums[i]) // 存储后半部分
	}
	// beego.Info(pre, post)
	// beego.Info("pre")
	for i := 0; i < len(post); i++ { // 重新填充前半部分
		nums[i] = post[i]
		// beego.Info(fmt.Sprintf("nums[%d]=%d", i, post[i]))
	}
	// beego.Info("post")
	for i := 0; i < len(pre); i++ { // 重新填充后半部分
		nums[k+i] = pre[i]
		// beego.Info(fmt.Sprintf("nums[%d]=%d", i, pre[i-k]))
	}
	// beego.Info(nums)
}
