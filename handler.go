package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	longtoshort "gocode/homework_shortUrl/longToShort"
	"gocode/homework_shortUrl/model"
	example "gocode/kitex_gen/kitex/example"
	"net/http"
)

// ShortUrlServiceImpl implements the last service interface defined in the IDL.
type ShortUrlServiceImpl struct{}

// Register implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) Register(ctx context.Context, request *example.RegisterRequest) (resp *example.RegisterResponse, err error) {
	// TODO: Your code here...
	if request.Name == "" || request.Passwd == "" {
		return nil, errors.New("注册信息不完整，无法继续进行")
	}
	result, err := model.DB.Exec("insert into users (name, passwd) value (?,?)",
		request.Name, request.Passwd)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("注册错误")
	}
	fmt.Println(result.LastInsertId())

	model.RDB.Set(ctx, "name", request.Name, 0)
	model.RDB.Set(ctx, "passwd", request.Passwd, 0)

	resp = &example.RegisterResponse{
		Status:  200,
		Message: fmt.Sprintf("用户 %s 注册成功", request.Name),
	}
	return resp, nil
}

// Login implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) Login(ctx context.Context, request *example.LoginRequest) (resp *example.LoginResponse, err error) {
	// TODO: Your code here...
	row := model.DB.QueryRow("select id from users where name=? and passwd=?",
		request.Name, request.Passwd)
	var id int
	if row != nil {
		err := row.Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("用户名或密码错误")
			} else {
				return nil, errors.New("服务器内部错误")
			}
		}
	}

	val, err := model.RDB.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("纯登录model.RDB.Get err=", err)
	}
	fmt.Println("name:", val)
	val, err = model.RDB.Get(ctx, "passwd").Result()
	if err != nil {
		fmt.Println("纯登录model.RDB.Get err=", err)
	}
	fmt.Println("passwd:", val)

	resp = &example.LoginResponse{
		Status:  200,
		Message: fmt.Sprintf("用户 %s 登录成功", request.Name),
	}
	return resp, nil
}

// WriteShortUrl implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) WriteShortUrl(ctx context.Context, request *example.ShortUrlRequest) (resp *example.ShortUrlResponse, err error) {
	// TODO: Your code here...
	row := model.DB.QueryRow("select id from users where long_url=?",
		request.LongUrl)
	var id int
	if row != nil {
		err := row.Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("long_url错误")
			} else {
				return nil, errors.New("服务器内部错误")
			}
		}
	}
	resp.ShortUrl = longtoshort.To(id, request.LongUrl)
	//更新短链接
	result, err := model.DB.Exec("update users set short_url=? where id=?",
		resp.ShortUrl, id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("短连接注入错误")
	}
	fmt.Println(result.LastInsertId())

	model.RDB.Set(ctx, resp.ShortUrl, request.LongUrl, 0) //以短链接为key,长链接为val，存入缓存

	resp = &example.ShortUrlResponse{
		ShortUrl: resp.ShortUrl,
	}
	return resp, nil
}

// Redirect implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) Redirect(ctx context.Context, request *example.RedirectRequest) (err error) {
	// TODO: Your code here...
	longUrl, err := model.RDB.Get(ctx, request.ShortUrl).Result() //从Redis缓存中获取长链接
	if err != nil {
		fmt.Println(err)
		return errors.New("short code not found")
	}
	var w http.ResponseWriter
	var r http.Request
	http.Redirect(w, &r, longUrl, http.StatusFound)
	return nil
}

// DeleteShortUrl implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) DeleteShortUrl(ctx context.Context, request *example.DeleteShortUrlRequest) (err error) {
	// TODO: Your code here...
	//查询短链接是否存在
	row := model.DB.QueryRow("select id from users where short_url=?",
		request.ShortUrl)
	var id int
	if row != nil {
		err = row.Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("短连接不存在")
			} else {
				return errors.New("服务器内部错误")
			}
		}
	}
	//删除短链接
	//更新mysql
	result, err := model.DB.Exec("update users set short_url=? where id=?",
		"", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.LastInsertId())
	//redis
	err = model.RDB.Del(ctx, "short_url").Err()
	if err != nil {
		fmt.Println("redis delete short_url err=", err)
	}

	return nil
}

// ChangeShortUrl implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) ChangeShortUrl(ctx context.Context, request *example.ChangeShortUrlRequest) (err error) {
	// TODO: Your code here...
	var longUrl string
	row := model.DB.QueryRow("select long_url from users where short_url=?",
		request.OldShortUrl)
	if row != nil {
		err = row.Scan(&longUrl)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("短码错误")
			} else {
				fmt.Println(err)
				return errors.New("服务器内部错误")
			}
		}
		if longUrl == "" || request.OldShortUrl == "" {
			return errors.New("长链接或短连接不存在")
		}
	}
	//更新短链接
	//mysql
	if model.DB != nil {
		result, err := model.DB.Exec("update users set short_url=? where long_url=?",
			request.NewShortUrl_, longUrl)
		if err != nil {
			fmt.Println(err)
			return errors.New("mysql更新自定义短链接错误")
		}
		fmt.Println(result.LastInsertId())
	} else {
		fmt.Println("空指针错误..")
	}
	//redis
	// 更新key
	err = model.RDB.Rename(ctx, request.OldShortUrl, request.NewShortUrl_).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

// ShowShortUrl implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) ShowShortUrl(ctx context.Context, request *example.ShowShortUrlRequest) (resp *example.ShowShortUrlResponse, err error) {
	// TODO: Your code here...
	row := model.DB.QueryRow("select short_url from users where name=?",
		request.Username)
	if row != nil {
		err = row.Scan(&resp.ShortUrl)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("用户名不存在")
			} else {
				return nil, errors.New("服务器内部错误")
			}
		}
	}
	return resp, nil
}

// RateShortUrl implements the ShortUrlServiceImpl interface.
func (s *ShortUrlServiceImpl) RateShortUrl(ctx context.Context) (resp *example.RateShortUrlResponse, err error) {
	// TODO: Your code here...
	// 查询同类所有数据
	rows, err := model.DB.Query("select short_url from users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var shortUrl string
	// 遍历结果集
	var data []string
	for rows.Next() {
		err := rows.Scan(&shortUrl)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, shortUrl)
	}

	// 输出结果
	for _, d := range data {
		fmt.Printf("%s\n", d)
	}

	resp = &example.RateShortUrlResponse{
		SortedShortUrls: data,
	}
	return resp, nil
}
