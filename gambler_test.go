package gambler

import "testing"

func TestGamblingSession_Attempts(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{ name: `simple`, fields: fields{attempts: 1}, want: 1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if got := gs.Attempts(); got != tt.want {
				t.Errorf("GamblingSession.Attempts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamblingSession_Wins(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{ name: `simple`, fields: fields{wins: 1}, want: 1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if got := gs.Wins(); got != tt.want {
				t.Errorf("GamblingSession.Wins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamblingSession_Losses(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{ name: `simple`, fields: fields{losses: 1}, want: 1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if got := gs.Losses(); got != tt.want {
				t.Errorf("GamblingSession.Losses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamblingSession_Balance(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{ name: `simple`, fields: fields{balance: 1}, want: 1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if got := gs.Balance(); got != tt.want {
				t.Errorf("GamblingSession.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamblingSession_BuyAttempts(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantAttempts int
		wantBalance int
	}{
		{ name: `simple`, fields: fields{attempts: 2, balance: 5}, args: args{quantity: 2}, wantErr: false, wantAttempts: 4, wantBalance: 3},
		{ name: `full balance`, fields: fields{attempts: 2, balance: 5}, args: args{quantity: 5}, wantErr: false, wantAttempts: 7, wantBalance: 0},
		{ name: `insufficient balance`, fields: fields{attempts: 2, balance: 5}, args: args{quantity: 6}, wantErr: true, wantAttempts: 2, wantBalance: 5},
		{ name: `zero quantity`, fields: fields{attempts: 2, balance: 5}, args: args{quantity: 0}, wantErr: true, wantAttempts: 2, wantBalance: 5},
		{ name: `negative quantity`, fields: fields{attempts: 2, balance: 5}, args: args{quantity: -3}, wantErr: true, wantAttempts: 2, wantBalance: 5},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if err := gs.BuyAttempts(tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("GamblingSession.BuyAttempts() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gs.attempts != tt.wantAttempts {
				t.Errorf("GamblingSession.attempts = %v, want %v", gs.attempts, tt.wantAttempts)
			}
			if gs.balance != tt.wantBalance {
				t.Errorf("GamblingSession.balance = %v, want %v", gs.balance, tt.wantBalance)
			}
		})
	}
}

func TestGamblingSession_MakeAnAttempt(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	type args struct {
		roll int
	}
	tests := []struct {
		name    string
		fields  fields
		args	args
		want    bool
		want1   int
		wantErr bool
		wantWins int
		wantLosses int
		wantAttempts int
	}{
		{ name: `simple loss`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: 5}, want: false , want1: 5, wantErr: false, wantWins: 2, wantLosses: 5, wantAttempts: 0},
		{ name: `simple win`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: 85}, want: true , want1: 85, wantErr: false, wantWins: 3, wantLosses: 4, wantAttempts: 0},
		{ name: `perfect win`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: 80}, want: true , want1: 80, wantErr: false, wantWins: 3, wantLosses: 4, wantAttempts: 0},
		{ name: `zero roll`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: 0}, want: false , want1: 0, wantErr: true, wantWins: 2, wantLosses: 4, wantAttempts: 1},
		{ name: `above expected range`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: 250}, want: false , want1: 250, wantErr: true, wantWins: 2, wantLosses: 4, wantAttempts: 1},
		{ name: `below expected range`, fields: fields{attempts: 1, wins: 2, losses: 4}, args: args{roll: -5}, want: false , want1: -5, wantErr: true, wantWins: 2, wantLosses: 4, wantAttempts: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			got, got1, err := gs.MakeAnAttempt(tt.args.roll)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamblingSession.MakeAnAttempt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GamblingSession.MakeAnAttempt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GamblingSession.MakeAnAttempt() got1 = %v, want %v", got1, tt.want1)
			}
			if gs.wins != tt.wantWins {
				t.Errorf("GamblingSession.attempts = %v, want %v", gs.wins, tt.wantWins)
			}
			if gs.losses != tt.wantLosses {
				t.Errorf("GamblingSession.attempts = %v, want %v", gs.losses, tt.wantLosses)
			}
			if gs.attempts != tt.wantAttempts {
				t.Errorf("GamblingSession.attempts = %v, want %v", gs.attempts, tt.wantAttempts)
			}
		})
	}
}

func TestGamblingSession_AddToBalance(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	type args struct {
		dollars int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantBalance int
	}{
		{ name: `simple`, fields: fields{balance: 5}, args: args{dollars: 2}, wantErr: false, wantBalance: 7},
		{ name: `adding negative`, fields: fields{balance: 5}, args: args{dollars: -4}, wantErr: true, wantBalance: 5},
		{ name: `adding zero`, fields: fields{balance: 5}, args: args{dollars: 0}, wantErr: true, wantBalance: 5},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if err := gs.AddToBalance(tt.args.dollars); (err != nil) != tt.wantErr {
				t.Errorf("GamblingSession.AddToBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gs.balance != tt.wantBalance {
				t.Errorf("GamblingSession.balance = %v, want %v", gs.balance, tt.wantBalance)
			}
		})
	}
}

func TestGamblingSession_CashOut(t *testing.T) {
	type fields struct {
		attempts int
		wins     int
		losses   int
		balance  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		wantAttempts int
		wantBalance int
	}{
		{ name: `simple`, fields: fields{attempts: 1, balance: 3}, wantErr: false, wantAttempts: 0, wantBalance: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GamblingSession{
				attempts: tt.fields.attempts,
				wins:     tt.fields.wins,
				losses:   tt.fields.losses,
				balance:  tt.fields.balance,
			}
			if err := gs.CashOut(); (err != nil) != tt.wantErr {
				t.Errorf("GamblingSession.CashOut() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gs.attempts != tt.wantAttempts {
				t.Errorf("GamblingSession.attempts = %v, want %v", gs.balance, tt.wantAttempts)
			}
			if gs.balance != tt.wantBalance {
				t.Errorf("GamblingSession.balance = %v, want %v", gs.balance, tt.wantBalance)
			}
		})
	}
}
