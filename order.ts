export interface Order {
  orderNo: string;
  version: number;
  orderCurrency: Currency;
  material: Material;
  changeDate: string;
  supplierNo: string;
  orderState: string;
  orderType: string;
  quantityUnit: string;
  priceUnit: string;
  pricingPeriods: PricingPeriod[];
}

export interface Material {
  snr: string;
  ai: number;
  plant: string;
}

export enum Currency {
  Eur = 'EUR',
}

export interface PricingPeriod {
  internalIdentifier: number;
  basePrice: BasePrice;
  nettPrice: BasePrice;
  validity: Validity;
  surcharges: Surcharge[];
  orderState: string;
  weightings: any[];
}

export interface BasePrice {
  currency: Currency;
  value: number;
}

export interface Surcharge {
  internalIdentifier: number;
  amount: BasePrice;
  category: string;
  surchargeType: string;
}

export interface Validity {
  fromDate: Date;
  toDate: Date;
}

export enum Plant{
  MSF='MSF',
  NEDCAR='NEDCAR'
}
