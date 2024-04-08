namespace go kitex.example

struct RegisterRequest {
  1: string name,
  2: string passwd,
}

struct RegisterResponse {
  1: i32 status,
  2: string message,
}

struct LoginRequest {
  1: string name,
  2: string passwd,
}

struct LoginResponse {
  1: i32 status,
  2: string message,
}

struct ShortUrlRequest {
  1: string longUrl,
}

struct ShortUrlResponse {
  1: string shortUrl,
}

struct RedirectRequest {
  1: string shortUrl,
}

struct DeleteShortUrlRequest {
  1: string shortUrl,
}

struct ChangeShortUrlRequest {
  1: string oldShortUrl,
  2: string newShortUrl,
}

struct ShowShortUrlRequest {
  1: string username,
}

struct ShowShortUrlResponse {
  1: string shortUrl,
}

struct RateShortUrlResponse {
  1: list<string> sortedShortUrls,
}

service ShortUrlService {
  RegisterResponse register(1: RegisterRequest request),
  LoginResponse login(1: LoginRequest request),
  ShortUrlResponse writeShortUrl(1: ShortUrlRequest request),
  void redirect(1: RedirectRequest request),
  void deleteShortUrl(1: DeleteShortUrlRequest request),
  void changeShortUrl(1: ChangeShortUrlRequest request),
  ShowShortUrlResponse showShortUrl(1: ShowShortUrlRequest request),
  RateShortUrlResponse rateShortUrl(),
}
