package destinyapi

type Vendor interface {
	Hash() string
	Name() string
}

type BountyTrackerVendor struct{}

func (BountyTrackerVendor) Hash() string { return "1527174714" }
func (BountyTrackerVendor) Name() string { return "Bounty Tracker" }

type Cayde6Vendor struct{}

func (Cayde6Vendor) Hash() string { return "3003633346" }
func (Cayde6Vendor) Name() string { return "Cayde-6" }

type CrucibleVendor struct{}

func (CrucibleVendor) Hash() string { return "3658200622" }
func (CrucibleVendor) Name() string { return "Crucible" }

type CryptarchVendor struct{}

func (CryptarchVendor) Hash() string { return "4269570979" }
func (CryptarchVendor) Name() string { return "Cryptarch" }

type DeadOrbitVendor struct{}

func (DeadOrbitVendor) Hash() string { return "3611686524" }
func (DeadOrbitVendor) Name() string { return "Dead Orbit" }

type EmblemKioskVendor struct{}

func (EmblemKioskVendor) Hash() string { return "3301500998" }
func (EmblemKioskVendor) Name() string { return "Emblem Kiosk" }

type EmoteKioskVendor struct{}

func (EmoteKioskVendor) Hash() string { return "614738178" }
func (EmoteKioskVendor) Name() string { return "Emote Kiosk" }

type ErisMornVendor struct{}

func (ErisMornVendor) Hash() string { return "174528503" }
func (ErisMornVendor) Name() string { return "Eris Morn" }

type EvaLevanteVendor struct{}

func (EvaLevanteVendor) Hash() string { return "134701236" }
func (EvaLevanteVendor) Name() string { return "Eva Levante" }

type EververseVendor struct{}

func (EververseVendor) Hash() string { return "3917130357" }
func (EververseVendor) Name() string { return "Eververse" }

type ExoticArmorKioskVendor struct{}

func (ExoticArmorKioskVendor) Hash() string { return "3902439767" }
func (ExoticArmorKioskVendor) Name() string { return "Exotic Armor Kiosk" }

type ExoticWeaponKioskVendor struct{}

func (ExoticWeaponKioskVendor) Hash() string { return "1460182514" }
func (ExoticWeaponKioskVendor) Name() string { return "Exotic Weapon Kiosk" }

type FutureWarCultVendor struct{}

func (FutureWarCultVendor) Hash() string { return "1821699360" }
func (FutureWarCultVendor) Name() string { return "Future War Cult" }

type GunsmithVendor struct{}

func (GunsmithVendor) Hash() string { return "570929315" }
func (GunsmithVendor) Name() string { return "Gunsmith" }

type IkoraReyVendor struct{}

func (IkoraReyVendor) Hash() string { return "1575820975" }
func (IkoraReyVendor) Name() string { return "Ikora Rey" }

type NewMonarchyVendor struct{}

func (NewMonarchyVendor) Hash() string { return "1808244981" }
func (NewMonarchyVendor) Name() string { return "New Monarchy" }

type PostmasterVendor struct{}

func (PostmasterVendor) Hash() string { return "2021251983" }
func (PostmasterVendor) Name() string { return "Postmaster" }

type ShaderKioskVendor struct{}

func (ShaderKioskVendor) Hash() string { return "2420628997" }
func (ShaderKioskVendor) Name() string { return "Shader Kiosk" }

type ShaxxVendor struct{}

func (ShaxxVendor) Hash() string { return "3746647075" }
func (ShaxxVendor) Name() string { return "Shaxx" }

type ShipKioskVendor struct{}

func (ShipKioskVendor) Hash() string { return "2244880194" }
func (ShipKioskVendor) Name() string { return "Ship Kiosk" }

type ShipwrightVendor struct{}

func (ShipwrightVendor) Hash() string { return "459708109" }
func (ShipwrightVendor) Name() string { return "Shipwright" }

type SparrowKioskVendor struct{}

func (SparrowKioskVendor) Hash() string { return "44395194" }
func (SparrowKioskVendor) Name() string { return "Sparrow Kiosk" }

type TheSpeakerVendor struct{}

func (TheSpeakerVendor) Hash() string { return "2680694281" }
func (TheSpeakerVendor) Name() string { return "The Speaker" }

type VanguardVendor struct{}

func (VanguardVendor) Hash() string { return "2668878854" }
func (VanguardVendor) Name() string { return "Vanguard" }

type ZavallaVendor struct{}

func (ZavallaVendor) Hash() string { return "1990950" }
func (ZavallaVendor) Name() string { return "Zavalla" }
